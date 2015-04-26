"use strict";

var React = require('react'),
    AppDispatcher = require('../lib/dispatcher'),
    constants = require('./constants');

/**
 * Splash
 *
 * Pre-login splash screen.
 */
var Splash = React.createClass({
    _buttonClick: function() {
        facebookConnectPlugin.login(['public_profile', 'email'],
                                    this._fbStatusChange.bind(this),
                                    console.log.bind(console));
    },
    _viewChange: function(view) {
        AppDispatcher.dispatch({
            actionType: constants.EXECUTIVE_ORDERS.CHANGE_VIEW,
            view: view
        });
    },
    _fbStatusChange: function(response) {
        console.log(response);
        if(response.status === 'connected') {
            this._viewChange(constants.VIEWS.MainMenu);
        }
    },
    componentWillMount: function() {
    },
    render: function() {
        return (
            <div>
                <div className="row">
                    <div className="col-xs-12">
                        <img className="img-responsive center-block"
                             src="http://ecx.images-amazon.com/images/I/51Nf7DSjczL._SY300_.png"></img>
                    </div>
                </div>
                <div className="row">
                    <div className="col-xs-6 col-xs-offset-3">
                        <button className="btn btn-default btn-block btn-lg"
                                onClick={this._buttonClick}>
                            Log In
                        </button>
                    </div>
                </div>
            </div>
        );
    }
});

module.exports = Splash;
