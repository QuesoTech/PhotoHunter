"use strict";

var React = require('react'),
    constants = require('./constants'),
    Judgebot9000 = require('./judgebot.jsx'),
    ResultViewer = require('./result_viewer.jsx'),
    AppDispatcher = require('../lib/dispatcher'),
    stores = require('../lib/stores');

/**
 * JusticeSystem
 *
 * Wrapper feeding image/label pairs to the Judgebot-9000. Handles submitting
 * data to the TestPairStore.
 */
var JusticeSystem = React.createClass({
    getInitialState: function() {
        return {
            testing: true,
            testPair: stores.TestPairStore.getNext(),
            results: null
        };
    },
    componentDidMount: function() {
        this._dispatcherId = AppDispatcher.register(this._onComplete);
    },
    componentWillUnmount: function() {
        AppDispatcher.unregister(this._dispatcherId);
    },
    render: function() {
        if(this.state.testing) {
            return <Judgebot9000 {...this.state.testPair} />;
        } else {
            return <ResultViewer results={this.state.results} />;
        }
    },
    _onComplete: function(action) {
        if(action.actionType == constants.ACTIONS.JUDGE_COMPLETE) {
            // Finished judging this pair, submit user choice and show results.
            stores.TestPairStore.submit(action.label);
            this.setState({
                testing: false,
                results: stores.TestPairStore.getResults()
            });
        } else if(action.actionType == constants.ACTIONS.RESULTS_DONE) {
            // finished showing results and not returning to menu, so show the
            // next test pair.
            this.setState({
                testing: true,
                testPair: stores.TestPairStore.getNext(),
                results: null
            });
        }
    }
});

module.exports = JusticeSystem;
