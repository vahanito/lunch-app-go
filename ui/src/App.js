import React, {Component} from 'react';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';
import './App.css';
import RestaurantList from 'pages/RestaurantList';
import Menu from 'pages/Menu';
import MenuTable from "./pages/MenuTable";
import {Loader} from "./components/Loader";

class App extends Component {
    render() {
        return (
            <Router>
                <>
                    <Switch>
                        <Route path="/" exact component={MenuTable}/>
                        <Route path="/restaurants" exact component={RestaurantList}/>
                        <Route path="/restaurants/:restaurant" exact component={Menu}/>
                        <Route path="/" render={() => (<div><p>404 Not found</p><Loader loading={true}/></div>)}/>
                    </Switch>
                </>
            </Router>
        );
    }
}

export default App;
