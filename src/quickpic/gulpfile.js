"use strict";

var gulp = require('gulp'),
    reactify = require('reactify'),
    browserify = require('browserify'),
    buffer = require('vinyl-buffer'),
    source = require('vinyl-source-stream'),
    sourcemaps = require('gulp-sourcemaps'),
    cdv = require('cordova-lib').cordova.raw,
    exec = require('child_process').exec;

var paths = {
    scripts: ["index.jsx", "lib/**/*.jsx", "views/**/*.jsx"]
};

var bundler = browserify({
    entries: ['./index.jsx'],
    transform: [reactify],
    debug: true
});

gulp.task('cordova:setup', function(done) {
    process.chdir('./cordova');
    exec('cordova platform add android@3.6.4', function(err) {
        exec('cordova -d plugins add ../plugins/phonegap-facebook-plugin --variable APP_ID=1662797127286769 --variable APP_NAME=QuickPic', done);
    });
});

gulp.task('cordova:teardown', function(done) {
    process.chdir('./cordova');
    exec('cordova platform rm android', function(err) {
        if(err) { done(err); } else {
            exec('cordova plugins rm com.phonegap.plugins.facebookconnect', done);
        }
    });
});

gulp.task('browserify', function() {
    return bundler.bundle()
        .pipe(source('quickpic.js'))
        .pipe(buffer())
        .pipe(sourcemaps.init({loadMaps: true}))
        .pipe(sourcemaps.write('./'))
        .pipe(gulp.dest('./cordova/www/js/dist'));
});

gulp.task('emulate', ['cordova:setup', 'browserify'], function(cb) {
    process.chdir('./cordova');
    cdv.emulate().then(cb);
});

gulp.task('serve', ['cordova:setup', 'browserify'], function(cb) {
    process.chdir('./cordova');
    cdv.serve().then(function() { cb(); });
});

gulp.task('deploy-local', ['cordova:setup', 'browserify'], function(cb) {
    process.chdir('./cordova');
    cdv.run().then(cb);
});

gulp.task('watch', function() {
    var tasks = process.argv.slice(3);
    return gulp.watch(paths.scripts, tasks);
});

gulp.task('default', ['deploy-local']);
