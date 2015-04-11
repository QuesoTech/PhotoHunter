"use strict";

var gulp = require('gulp'),
    reactify = require('reactify'),
    browserify = require('browserify'),
    buffer = require('vinyl-buffer'),
    source = require('vinyl-source-stream'),
    sourcemaps = require('gulp-sourcemaps'),
    cdv = require('cordova-lib').cordova.raw;

var paths = {
    scripts: ["index.jsx", "lib/**/*.jsx", "views/**/*.jsx"]
};

var bundler = browserify({
    entries: ['./index.jsx'],
    transform: [reactify],
    debug: true
});

gulp.task('browserify', function() {
    return bundler.bundle()
        .pipe(source('quickpic.js'))
        .pipe(buffer())
        .pipe(sourcemaps.init({loadMaps: true}))
        .pipe(sourcemaps.write('./'))
        .pipe(gulp.dest('./cordova/www/js/dist'));
});

gulp.task('emulate', ['browserify'], function(cb) {
    process.chdir('./cordova');
    cdv.emulate().then(cb);
});

gulp.task('serve', ['browserify'], function(cb) {
    process.chdir('./cordova');
    cdv.serve().then(function() { cb(); });
});

gulp.task('watch', function() {
    var tasks = process.argv.slice(3);
    return gulp.watch(paths.scripts, tasks);
});

gulp.task('default', ['browserify']);
