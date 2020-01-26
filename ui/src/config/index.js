import local from './local';
import prod from './prod';

let config = local;

if (process.env.REACT_APP_ENV === 'prod') {
    config = prod;
}

export default {
    ...config
};
