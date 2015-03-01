"use strict";
/** @jsx React.createElement */

var React = require('react');

/**
 * Flasher
 *
 * The *display* of an image flashing up. Does *not* handle the flashing
 * itself. That is handled by the Judge-Bot 9000.
 */
var Flasher = React.createClass({
    render: function() {
        return <img className="img-responsive center-block" {...this.props}></img>;
    }
});

/**
 * Label
 *
 * A label to be chosen from a LabelList.
 */
var Label = React.createClass({
    propTypes: {
        value: React.PropTypes.string.isRequired,
        onClick: React.PropTypes.func,
    },
    getDefaultProps: function() {
        return {
            onClick: function() {}
        };
    },
    render: function() {
        return (<li className="btn btn-default btn-block btn-lg" onClick={this.props.onClick}>{this.props.value}</li>);
    }
});

/**
 * LabelList
 *
 * A list of labels.
 */
var LabelList = React.createClass({
    propTypes: {
        labels: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
        onChoice: React.PropTypes.func,
    },
    getDefaultProps: function() {
        return {
            onChoice: function() {}
        };
    },
    render: function() {
        var labels = this.props.labels.map(function(label, index) {
            return <Label value={label} key={index} onClick={this.props.onChoice.bind(null, index)} />;
        }.bind(this));

        return (<ul className="list-unstyled">{labels}</ul>);
    }
});

/**
 * The Judge-Bot 9000
 *
 * First, show an image for N (default: 5) seconds. Then, display a list of
 * possible labels. Finally, inform the user of correctness (?) and prep to
 * move to the next image.
 */
var JudgeBot9000 = React.createClass({
    propTypes: {
        src: React.PropTypes.string.isRequired,
        labels: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
        time: React.PropTypes.number,
        onComplete: React.PropTypes.func
    },
    getDefaultProps: function() {
        return {
            time: 5000,
            onComplete: function() { console.log("done"); }
        };
    },
    getInitialState: function() {
        return {
            flashing: true,
            timeoutID: null
        };
    },
    onImageLoad: function() {
        if(this.state.timeoutID === null) {
            var timeoutID = setTimeout(function() {
                this.setState({flashing: false});
            }.bind(this), this.props.time);

            this.setState({timeoutID: timeoutID});
        }
    },
    onChooseLabel: function() {
        var that = this;
        return function(index) {
            // submitLabel(...)
            that.props.onComplete();
        };
    },
    render: function() {
        var content;
        if(this.state.flashing) {
            content = <Flasher src={this.props.src} onLoad={this.onImageLoad}/>;
        } else {
            content = <LabelList labels={this.props.labels} onChoice={this.onChooseLabel()}/>;
        }

        return (
            <div className="row">
                <div className="col-xs-12">
                    <h3 className="text-center">{"What kind of object is in this scene?"}</h3>
                </div>
                <div className="col-xs-12">
                    {content}
                </div>
            </div>
        );
    }
});

module.exports = JudgeBot9000;
