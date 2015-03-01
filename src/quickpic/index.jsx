"use strict";
/** @jsx React.createComponent */

var JudgeBot9000 = require('./lib/judgebot.jsx'),
    React = require('react');

function onload() {
    React.render(
        <JudgeBot9000 src='http://o.aolcdn.com/hss/storage/adam/dde819906acab479c31e42b7bf468281/fear%20bot%20planet%201%20mac%20judging.jpg'
                      time={1000}
                      labels={['Robot','Judge','Cat', 'Tree']} />,
        document.querySelector('#react-container')
    );
}

if(document.hasOwnProperty('ondeviceready')) {
    document.addEventListener('deviceready', onload);
} else {
    document.addEventListener('DOMContentLoaded', onload);
}
