var browserify = require('browserify');
var gulp = require('gulp');
var handleErrors = require('../util/handleErrors');
var source = require('vinyl-source-stream');

gulp.task('browserify', function(){
  return browserify({
    entries: ['./src/js/game.js'],
  })
  .bundle({debug: true})
  .on('error', handleErrors)
  .pipe(source('game.js'))
  .pipe(gulp.dest('./public/js/'));
});
