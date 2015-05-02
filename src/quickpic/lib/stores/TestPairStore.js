"use strict";

var AppDispatcher = require('../dispatcher'),
    EventEmitter = require('events').EventEmitter,
    assign = require('object-assign');

var _current = {
    _id: 42,
    src: 'http://o.aolcdn.com/hss/storage/adam/dde819906acab479c31e42b7bf468281/fear%20bot%20planet%201%20mac%20judging.jpg',
    labels: ['Robot','Judge','Cat', 'Tree']
};

/**
 * TestPairStore
 *
 * A datastore communicating with the PhotoHunter API which mediates retrieval
 * of Image/Label test pairs and the submission of label choices.
 */
var TestPairStore = assign({}, EventEmitter.prototype, {
    /**
     * getNext
     *
     * Retrieve next image/label test pair.
     */
    getNext: function() {
        return _current;
    },
    /**
     * submit
     *
     * Submit a label for the current test pair. Callable once per test pair
     * only.
     */
    submit: function(label) {
        // do nothing until we have an api
    },
    /**
     * getResults
     *
     * Retrieves results from the API for the current test pair.
     */
    getResults: function() {
        return {
            votes: {
                'Robot': 9001,
                'Judge': 1009,
                'Cat':   10,
                'Tree':  0
            }
        };
    }
});

module.exports = TestPairStore;
