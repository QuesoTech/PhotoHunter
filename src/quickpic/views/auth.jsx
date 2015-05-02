"use strict";

var React = require('react'),
    AppDispatcher = require('../lib/dispatcher'),
    constants = require('./constants');

/**
 * Authentication
 *
 * Docstring goes here.
 */
var Authentication = React.createClass({
    _moveAlong: function() {
        AppDispatcher.dispatch({
            actionType: constants.EXECUTIVE_ORDERS.CHANGE_VIEW,
            view: constants.VIEWS.MainMenu
        });
    },
    render: function() {
        return (
            <div>
                <div className="row">
                    <div className="col-xs-12">
                        <h3 className="text-center">Move along, nothing to see here.</h3>
                    </div>
                </div>
                <div className="row">
                    <div className="col-xs-12">
                        <button className="btn btn-default btn-block btn-lg" onClick={this._moveAlong}>
                            Yes, Ociffer
                        </button>
                    </div>
                </div>
            </div>
        );
    }
});

module.exports = Authentication;
