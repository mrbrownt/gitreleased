// vue.config.js

module.exports = {
    devServer: {
        proxy: {
            "^/api": {
                target: "http://web:3000",
            },
            "^/auth": {
                target: "http://web:3000",
            },
        },
    },
}
