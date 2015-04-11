"use strict";
/** @jsx React.createElement */

var React = require('react'),
    AppDispatcher = require('../lib/dispatcher'),
    constants = require('./constants');

var MenuButton = React.createClass({
    propTypes: {
        value: React.PropTypes.string.isRequired,
        onClick: React.PropTypes.func.isRequired
    },
    render: function() {
        return (
            <li className="btn btn-default btn-block btn-lg" onClick={this.props.onClick}>{this.props.value}</li>
        );
    }
});

var MainMenu = React.createClass({
    _buttonClick: function(which) {
        return function() {
            AppDispatcher.dispatch({
                actionType: constants.EXECUTIVE_ORDERS.CHANGE_VIEW,
                view: which
            });
        }
    },
    render: function() {
        return (
            <ul className="list-unstyled">
                <MenuButton value="Photo Quiz"
                            onClick={this._buttonClick(constants.VIEWS.Judgement)} />
                <MenuButton value="Statistics"
                            onClick={this._buttonClick(constants.VIEWS.Statistics)} />
                <MenuButton value="Leaderboard"
                            onClick={this._buttonClick(constants.VIEWS.Leaderboard)} />
            </ul>
        );
    }
});

module.exports = MainMenu;
