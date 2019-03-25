const { gulp, src, dest, series, parallel, watch } = require('gulp');
const autoprefixer = require('autoprefixer');
const cssnano      = require('cssnano');
const del          = require('del');
const postcss      = require('gulp-postcss');
const rename       = require('gulp-rename');
const sass         = require('gulp-sass');
const browserSync  = require('browser-sync').create();

function clean(cb) {
    del(['./build/']);
    cb();
}

function css() {
    var processors = [
		autoprefixer({ grid: "autoplace", browsers: ['last 3 versions', 'ie 8', 'ie 9', 'ie 10'] }),
		cssnano(),
	];
    return src('./assets/scss/**/*.scss')
        .pipe(sass({ includePaths: ['node_modules/'], outputStyle: 'expanded' }))
        .pipe(dest('./assets/css/'))
        .pipe(rename({ suffix: '.min' }))
        .pipe(postcss(processors))
        .pipe(dest('./assets/css/'))
        .pipe(browserSync.stream());
}

function build(cb) {
  css();
  cb();
}

function watchFiles() {
    browserSync.init({
        proxy: 'http://localhost:9071',
    });

    watch( './assets/scss/**/*.scss', series( css ) );
    watch( './templates/*.html', series( browserSync.reload ) );
}

const dev = series( build, watchFiles );

exports.watch = watchFiles;
exports.build = build;

exports.default = dev;

