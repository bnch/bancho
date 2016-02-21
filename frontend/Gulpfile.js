var gulp = require('gulp');
var concat = require('gulp-concat');
var cssnano = require('gulp-cssnano');
var autoprefixer = require('gulp-autoprefixer');
var uglify = require('gulp-uglify');
var rename = require('gulp-rename');

gulp.task('css', function(){
	gulp.src('css/**/*.css')
		.pipe(cssnano())
		.pipe(autoprefixer('last 2 version', 'safari 5', 'ie 8', 'ie 9'))
		.pipe(concat('style.min.css'))
		.pipe(gulp.dest('static'))
});
gulp.task('js', function() {
	gulp.src(['javascript/**/*.js', 'javascript/*.js'])
		.pipe(concat('script.js'))
		.pipe(gulp.dest('static'))
		.pipe(rename('script.min.js'))
		.pipe(uglify())
		.pipe(gulp.dest('static'))
});
gulp.task('default', ['css', 'js'], function(){});
