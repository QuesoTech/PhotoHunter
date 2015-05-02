"use strict";
/** @jsx React.createComponent */

var QuickPic = require('./views/index.jsx'),
    React = require('react');

function onload() {
    React.render(
        <QuickPic />,
        document.querySelector('#react-container')
    );
}

if(document.hasOwnProperty('ondeviceready')) {
    document.addEventListener('deviceready', onload);
} else {
    document.addEventListener('DOMContentLoaded', onload);
}
