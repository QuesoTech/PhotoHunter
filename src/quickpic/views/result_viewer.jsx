"use strict";

var React = require('react'),
    constants = require('./constants'),
    AppDispatcher = require('../lib/dispatcher'),
    _ = require('lodash');

var ResultChart = React.createClass({
    propTypes: {
        results: React.PropTypes.object.isRequired
    },
    componentDidMount: function() {
        var data = _(this.props.results.votes)
                         .map(function(val, key) {
                             return { label: key, value: val };
                         })
                         .value();

        var chart_container = document.querySelector('#result-chart');
        var chart = d3.select('#result-chart')
                      .append('svg')
                      .attr('width', chart_container.offsetWidth)
                      .attr('height', chart_container.offsetHeight)
                      .chart('Pie', {
                          width: chart_container.offsetWidth * 0.9,
                          height: chart_container.offsetHeight * 0.9,
                          radius: Math.min(chart_container.offsetWidth,
                                           chart_container.offsetHeight) * 0.4
                      });
        chart.draw(data);
    },
    render: function() {
        return <div id="result-chart" style={{height: "80%"}} />;
    }
});

var ResultFooter = React.createClass({
    render: function() {
        return (
            <div className="btn-group col-xs-12" style={{"margin-top": "10%"}} role="group">
                <button type="button"
                        className="btn col-xs-6 btn-info btn-lg"
                        role="button"
                        onClick={this._onMenu}>Menu</button>
                <button type="button"
                        className="btn col-xs-6 btn-success btn-lg"
                        role="button"
                        onClick={this._onDone}>Continue</button>
            </div>
        );
    },
    _onDone: function() {
        AppDispatcher.dispatch({
            actionType: constants.ACTIONS.RESULTS_DONE
        });
    },
    _onMenu: function() {
        AppDispatcher.dispatch({
            actionType: constants.EXECUTIVE_ORDERS.CHANGE_VIEW,
            view: constants.VIEWS.MainMenu
        });
    }
});

var ResultViewer = React.createClass({
    propTypes: {
        results: React.PropTypes.object.isRequired
    },
    render: function() {
        return (
            <div style={{height: "100%"}}>
                <ResultChart results={this.props.results} />
                <ResultFooter />
            </div>
        );
    }
});

module.exports = ResultViewer;
