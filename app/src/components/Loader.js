import {css} from "@emotion/core";
import {PacmanLoader} from "react-spinners";
import React from "react";

const override = css`
    position: fixed !important;
    top: 50%;
    left: 50%;
    margin-top: -50px; 
    margin-left: -80px;
`;

export function Loader(props) {
    return (
        <div className="loader-wrapper" hidden={!props.loading}>
            <PacmanLoader
                css={override}
                sizeUnit={"px"}
                size={40}
                color={'#007bff'}
                loading={props.loading}
            />
        </div>);
}
