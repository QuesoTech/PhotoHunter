"use strict";

var React = require('react'),
    AppDispatcher = require('../lib/dispatcher'),
    MainMenu = require('./main_menu.jsx'),
    JusticeSystem = require('./justice_system.jsx'),
    VIEWS = require('./constants').VIEWS,
    EXECUTIVE_ORDERS = require('./constants').EXECUTIVE_ORDERS;

var QuickPic = React.createClass({
    _viewControl: function(action) {
        if(action.actionType == EXECUTIVE_ORDERS.CHANGE_VIEW) {
            this.setState({view: action.view});
        }
    },
    getInitialState: function() {
        return {
            view: VIEWS.MainMenu
        };
    },
    componentDidMount: function() {
        this._id = AppDispatcher.register(this._viewControl);
    },
    componentWillUnmount: function() {
        AppDispatcher.unregister(this._id);
    },
    render: function() {
        switch(this.state.view) {
            case VIEWS.Authentication:
                return null;
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
