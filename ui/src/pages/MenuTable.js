import React, {Component} from "react";
import MenuCard from "./menu/MenuCard";
import CardGroup from "react-bootstrap/CardGroup";
import {SearchInput} from "../components/Input";
import {Loader} from "../components/Loader";
import { getMenu, getRestaurants } from '../api/Api';

class MenuTable extends Component {

    constructor(props) {
        super(props);
        this.state = {
            loading: false,
            menus: []
        }
    }

    componentDidMount() {
        this.setState({loading: true});
        getRestaurants().then(response => {
            response.forEach(restaurant => {
              getMenu(restaurant).then(menuResponse => {
                this.setState({
                  loading: false,
                  menus: this.state.menus.concat(menuResponse)
                })
              })
            })
        }).catch(reason => {
            this.setState({
                hasError: true,
                reason: reason
            });
        });
    }

    onInputChange(value) {
        this.setState(prevState => ({
            hasError: prevState.hasError,
            loading: prevState.loading,
            menus: prevState.menus.map(menu => {
                let searchString = value.toLowerCase();
                if (menu.soup && menu.soup.name) {
                    menu.highlightSoup = !!searchString && menu.soup.name.toLowerCase().includes(searchString);
                }
                menu.hidden = !menu.restaurant.toLowerCase().includes(searchString) && !menu.highlightSoup;
                menu.highlightMenus = [];
                menu.menus.forEach((menuitem, index) => {
                    if (!!searchString && menuitem.name.toLowerCase().includes(searchString)) {
                        menu.highlightMenus.push(index);
                        menu.hidden = false;
                    }
                });
                return menu;
            }),
        }));
    }

    render() {
        const restaurantsRow = this.state.menus.map(menu => {
            return !menu.hidden && <MenuCard menu={menu}/>
        });
        if (this.state.hasError) {
            return (
                <div>
                    <h2>Chyba</h2>
                </div>
            )
        }

        return (
            <div className="container-fluid">
				<div className="background" />
                <SearchInput onChange={this.onInputChange.bind(this)}/>
                <div className="row">
                    <CardGroup bsPrefix="cardGroup-modified">
                        {restaurantsRow}
                    </CardGroup>
                </div>
                <Loader loading={this.state.loading}/>
            </div>
        );
    }
}

export default MenuTable;
