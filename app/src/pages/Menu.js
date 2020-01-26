import React, {Component} from 'react';
import {getMenu} from "api/Api";
import MenuCard from "./menu/MenuCard";
import {Loader} from "../components/Loader";

class Menu extends Component {


    constructor(props) {
        super(props);
        this.restaurant = this.props.match.params.restaurant;
        this.state = {
            loading: false,
            menu: {
                soup: '',
                menus: []
            }
        }
    }

    componentDidMount() {
        this.setState({loading: true})
        getMenu(this.restaurant).then(response => {
            this.setState({
                menu: response,
                loading: false
            })
        }).catch(reason => {
            this.setState({hasError: true});
        });
    }

    render() {
        if (this.state.hasError) {
            return (<h2>Chyba</h2>)
        }
        return (
            <div>
                <MenuCard menu={this.state.menu}/>
                <Loader loading={this.state.loading}/>
            </div>
        );
    }

}

export default Menu;
