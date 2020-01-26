import config from 'config/';

const request = (options) => {
    const headers = new Headers({
        'Content-Type': 'application/json',
        'Accept': 'application/json'
    });

    // if(localStorage.getItem(ACCESS_TOKEN)) {
    //     headers.append('Authorization', 'Bearer ' + localStorage.getItem(ACCESS_TOKEN))
    // }

    const defaults = {headers: headers};
    options = Object.assign({}, defaults, options);

    return fetch(options.url, options)
        .then(response => {
            if (!response.ok) {
                return Promise.reject(response);
            }
            return response;
        })
        .then(response => response.json());
};

export function getRestaurants() {
    return request({
        url: "/api/restaurants",
        method: 'GET'
    });
}

export function getMenu(restaurant) {
    return request({
        url: "/api/restaurants/" + restaurant,
        method: 'GET'
    });
}


export function getMenus() {
    return request({
        url: "/api/menus",
        method: 'GET'
    });
}
