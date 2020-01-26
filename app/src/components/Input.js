import InputGroup from "react-bootstrap/InputGroup";
import React from "react";
import {FormControl} from "react-bootstrap";

export function SearchInput(props) {
    return (
        <div className="row" style={{paddingTop: '10px'}}>
            <div className="col-md-6">
                <InputGroup className="mb-3">
                    <InputGroup.Prepend>
                        <InputGroup.Text id="basic-addon1">Vyhľadaj</InputGroup.Text>
                    </InputGroup.Prepend>
                    <FormControl
                        placeholder="Zadaj názov reštaurácie alebo jedla"
                        aria-label="search"
                        aria-describedby="basic-addon1"
                        onChange={
                            event => {
                                props.onChange(event.target.value)
                            }
                        }
                    />
                </InputGroup>
            </div>
        </div>);
}
