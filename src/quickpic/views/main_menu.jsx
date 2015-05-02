"use strict";
/** @jsx React.createElement */

var React = require('react'),
    AppDispatcher = require('../lib/dispatcher'),
    constants = require('./constants');

/**
 * MenuButton
 *
 * Simple component representing a button on the menu screen.
 */
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

/**
 * MainMenu
 *
 * Component displaying the main menu. Dispatches requests to change view on
 * button clicks, which are handled in the main QuickPic component.
 */
var MainMenu = React.createClass({
    _buttonClick: function(which) {
        return function() {
            AppDispatcher.dispatch({
                actionType: constants.EXECUTIVE_ORDERS.CHANGE_VIEW,
                view: which
            });
        };
    },
    render: function() {
        return (
            <div>
                <div className="row">
                    <div className="col-xs-12">
                        <h3 className="text-center">QuickPic</h3>
                    </div>
                </div>
                <div className="row">
                    <div className="col-xs-12">
                        <ul className="list-unstyled">
                            <MenuButton value="Photo Quiz"
                                        onClick={this._buttonClick(constants.VIEWS.Judgement)} />
                            <MenuButton value="Statistics"
                                        onClick={this._buttonClick(constants.VIEWS.Statistics)} />
                            <MenuButton value="Leaderboard"
                                        onClick={this._buttonClick(constants.VIEWS.Leaderboard)} />
                        </ul>
                    </div>
                </div>
            </div>
        );
    }
});

module.exports = MainMenu;
