"use strict";

var React = require('react'),
    AppDispatcher = require('../lib/dispatcher'),
    MainMenu = require('./main_menu.jsx'),
    JusticeSystem = require('./justice_system.jsx'),
    Splash = require('./splash.jsx'),
    Authentication = require('./auth.jsx'),
    VIEWS = require('./constants').VIEWS,
    EXECUTIVE_ORDERS = require('./constants').EXECUTIVE_ORDERS;

/**
 * QuickPic
 *
 * Main component of the QuickPic application. Handles switching view
 * segments. Control returns here when any sub-component ceases control.
 */
var QuickPic = React.createClass({
    _viewControl: function(action) {
        if(action.actionType == EXECUTIVE_ORDERS.CHANGE_VIEW) {
            this.setState({view: action.view});
        }
    },
    getInitialState: function() {
        return {
            view: VIEWS.Splash
        };
    },
    componentDidMount: function() {
        this._id = AppDispatcher.register(this._viewControl);
    },
    componentWillUnmount: function() {
        AppDispatcher.unregister(this._id);
    },
    render: function() {
        // null ⇒ not yet implemented
        switch(this.state.view) {
            case VIEWS.Splash:
                return <Splash />;
            case VIEWS.Authentication:
                return <Authentication />;
            case VIEWS.MainMenu:
                return <MainMenu />;
            case VIEWS.Judgement:
                return <JusticeSystem />;
            case VIEWS.Statistics:
                return null;
            case VIEWS.Leaderboard:
                return null;
        }
    }
});

module.exports = QuickPic;
