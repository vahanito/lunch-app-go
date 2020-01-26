import React, {Component} from 'react';
import {getRestaurants} from 'api/Api';
import List from './restaurantList/List';

class RestaurantList extends Component {

    constructor(props) {
        super(props);
        this.state = {
            restaurants: []
        };
    }

    componentDidMount() {
        getRestaurants().then(response => {
            this.setState({restaurants: response});
        }).catch(reason => {
            this.setState({hasError: true});
        });
    }

    render() {
        if (this.state.hasError) {
            return (<h2>Chyba</h2>)
        }
        return (<List restaurants={this.state.restaurants}/>);
    }
}

export default RestaurantList;
