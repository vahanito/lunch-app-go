import React from 'react';
import PropTypes from 'prop-types';
import ListGroup from 'react-bootstrap/ListGroup';
import {Link} from "react-router-dom";

function List(props) {
    const restaurantList = props.restaurants.map(restaurant =>
        <ListGroup.Item key={restaurant}>
            <Link to={"/restaurants/" + restaurant}>{restaurant}</Link>
        </ListGroup.Item>);
    return (
        <ListGroup>
            <ListGroup.Item key="all">
                <Link to="/">Všetky</Link>
            </ListGroup.Item>
            {restaurantList}
        </ListGroup>
    );
}

List.propTypes = {
    restaurants: PropTypes.array.isRequired,
};

export default List;
