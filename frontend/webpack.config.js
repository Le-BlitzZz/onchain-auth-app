const path = require("path");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const ESLintPlugin = require("eslint-webpack-plugin");
const { WebpackManifestPlugin } = require("webpack-manifest-plugin");
const webpack = require("webpack");
const isDev = process.env?.BUILD_ENV === "development" || process.env?.NODE_ENV === "development";
const appName = "AuthOnchain";
const { VueLoaderPlugin } = require("vue-loader");
const { VuetifyPlugin } = require("webpack-plugin-vuetify");
const { DefinePlugin } = require("webpack");

const PATHS = {
  src: path.join(__dirname, "src"),
  css: path.join(__dirname, "src/css"),
  modules: path.join(__dirname, "node_modules"),
  app: path.join(__dirname, "src/app.js"),
  build: path.join(__dirname, "../assets/static/build"),
};

if (isDev) {
  console.log(`Starting ${appName} DEVELOPMENT build. Please wait.`);
} else {
  console.log(`Starting ${appName} PRODUCTION build. Please wait.`);
}

const config = {
  mode: isDev ? "development" : "production",
  devtool: isDev ? "inline-source-map" : false,
  optimization: {
    minimize: !isDev,
  },
  entry: {
    app: PATHS.app,
  },
  output: {
    path: PATHS.build,
    publicPath: "auto",
    filename: "[name].[contenthash].js",
    chunkFilename: "chunk/[name].[contenthash].js",
    asyncChunks: true,
    clean: true,
  },
  resolve: {
    modules: [PATHS.src, PATHS.modules],
    preferRelative: true,
    alias: {
      "vue$": "vue/dist/vue.runtime.esm-bundler.js",
    },
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: "[name].[contenthash].css",
      experimentalUseImportModule: false,
    }),
    new WebpackManifestPlugin({
      fileName: "assets.json",
      publicPath: "",
    }),
    new webpack.ProgressPlugin(),
    new VueLoaderPlugin(),
    new VuetifyPlugin({ autoImport: true }),
    new DefinePlugin({
      __VUE_OPTIONS_API__: JSON.stringify(true), // Change to false as needed
      __VUE_PROD_DEVTOOLS__: JSON.stringify(false), // Change to true to enable in production
      __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: JSON.stringify(false), // Change to true for detailed warnings
    }),
  ],
  performance: {
    hints: isDev ? false : "warning",
    maxEntrypointSize: 7500000,
    maxAssetSize: 7500000,
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: "vue-loader",
        options: {
          loaders: {
            js: "babel-loader",
            css: "css-loader",
          },
          compilerOptions: {
            whitespace: "preserve",
          },
        },
      },
      {
        test: /\.js$/,
        include: PATHS.src,
        exclude: (file) => /node_modules/.test(file),
        use: [
          {
            loader: "babel-loader",
            options: {
              sourceMap: isDev,
              compact: false,
              presets: ["@babel/preset-env"],
              plugins: [],
            },
          },
        ],
      },
      {
        test: /\.json$/,
        include: PATHS.src,
        type: "json",
      },
      {
        test: /\.css$/,
        include: [PATHS.css],
        exclude: /node_modules/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
          },
          {
            loader: "css-loader",
            options: {
              sourceMap: true,
              importLoaders: 1,
            },
          },
          "resolve-url-loader",
          {
            loader: "postcss-loader",
            options: {
              sourceMap: true,
              postcssOptions: {
                config: path.resolve(__dirname, "./postcss.config.js"),
              },
            },
          },
        ],
      },
      {
        test: /\.css$/,
        include: /node_modules/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
          },
          {
            loader: "css-loader",
            options: {
              sourceMap: true,
              importLoaders: 1,
            },
          },
          "resolve-url-loader",
          {
            loader: "postcss-loader",
            options: {
              sourceMap: true,
              postcssOptions: {
                config: path.resolve(__dirname, "./postcss.config.js"),
              },
            },
          },
        ],
      },
      {
        test: /\.s[c|a]ss$/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
          },
          {
            loader: "css-loader",
            options: {
              sourceMap: true,
              importLoaders: 1,
            },
          },
          "resolve-url-loader",
          {
            loader: "postcss-loader",
            options: {
              sourceMap: true,
              postcssOptions: {
                config: path.resolve(__dirname, "./postcss.config.js"),
              },
            },
          },
          "sass-loader",
        ],
      },
      {
        test: /\.(png|jpg|jpeg|gif)$/,
        type: "asset/resource",
        dependency: { not: ["url"] },
      },
      {
        test: /\.(woff(2)?|ttf|eot)(\?v=\d+\.\d+\.\d+)?$/,
        type: "asset/resource",
        dependency: { not: ["url"] },
      },
      {
        test: /\.svg/,
        type: "asset/resource",
        dependency: { not: ["url"] },
      },
    ],
  },
};

// Don't create sourcemap for production builds.
if (isDev) {
  const devToolPlugin = new webpack.SourceMapDevToolPlugin({
    filename: "[file].map",
  });

  config.plugins.push(devToolPlugin);

  import("eslint-formatter-pretty").then(() => {
    const esLintPlugin = new ESLintPlugin({
      formatter: "eslint-formatter-pretty",
      extensions: ["js"],
    });
    config.plugins.push(esLintPlugin);
  });
}

module.exports = config;
