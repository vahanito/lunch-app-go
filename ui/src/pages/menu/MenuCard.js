import React, {Component} from 'react'
import Table from 'react-bootstrap/Table'
import Card from 'react-bootstrap/Card'
import PropTypes from 'prop-types'
import Collapse from 'react-bootstrap/Collapse'

class MenuCard extends Component {

    constructor(props) {
        super(props);

		if (!localStorage.getItem(this.props.menu.restaurant)) {
			localStorage.setItem(this.props.menu.restaurant, 'true');
		}
        this.state = {
            open: localStorage.getItem(this.props.menu.restaurant) === 'true'
        }
    }

    render() {
        const tableBody = this.props.menu.soup ? this.menuList()
            : <tbody>
            <tr>
                <td>{this.props.menu.info ? this.props.menu.info : 'Chyba pri načítaní menu'}</td>
            </tr>
            </tbody>;
        return (
            <Card border="primary" bsPrefix="card card-no-break">
                <Card.Body bsPrefix="card-body card-body-less-padding">
                    <div className="card-header-wrapper card-header-padding">
                        <Card.Title>
                            <a href={this.props.menu.url} target="_blank" rel="noopener noreferrer">
                                {this.props.menu.restaurant} {
                                // this.props.menu.date ?
                                // <Moment parse="YYYY-MM-DD" format="DD.MM.YYYY">
                                // 	{this.props.menu.date}
                                // </Moment>
                                // : ""
                            }
                            </a>
                        </Card.Title>
                        {this.renderCollapseButton()}
                    </div>
                    <Collapse in={this.state.open}>
                        <div className="col-sm-12 table-collapse">
                            <Table>
                                {tableBody}
                            </Table>
                        </div>
                    </Collapse>
                </Card.Body>
            </Card>);
    }

    open = () => {
		localStorage.setItem(this.props.menu.restaurant, 'true');
		this.setState({open: true});
	}

	close = () => {
		localStorage.setItem(this.props.menu.restaurant, 'false');
		this.setState({open: false});
	}

    renderCollapseButton() {
        return (
            <div className="collapse-wrapper">
                <i className="fas fa-caret-down"
                   onClick={this.open}
                   aria-controls="table-collapse"
                   aria-expanded={this.state.open}
                   hidden={this.state.open}/>
                <i className="fas fa-caret-up"
                   onClick={this.close}
                   aria-controls="table-collapse"
                   aria-expanded={this.state.open}
                   hidden={!this.state.open}/>
            </div>
        )
    }

    menuList() {
        const menuList = this.props.menu.menus.map((menu, index) =>
            <tr key={index}>
                <td className="no-wrap"><i className="fas fa-utensils"></i> {index + 1}</td>
                <td className={this.highlightMenu(index)}>{menu.name}</td>
                {menu.price ? <td className="no-wrap">{menu.price}</td> : ""}
            </tr>);
        return <tbody>
        <tr>
            <td><i class="fas fa-utensil-spoon"></i></td>
            <td className={this.props.menu.highlightSoup ? 'font-weight-bold' : null}>{this.props.menu.soup.name}</td>
        </tr>
        {menuList}
        </tbody>;
    }

    highlightMenu(index) {
        return this.props.menu.highlightMenus && this.props.menu.highlightMenus.includes(index) ? 'font-weight-bold' : null;
    }
}

MenuCard.propTypes = {
    menu: PropTypes.exact({
        restaurant: PropTypes.string,
        date: PropTypes.any,
        soup: PropTypes.any,
        menus: PropTypes.array
    }).isRequired,
};

export default MenuCard;
