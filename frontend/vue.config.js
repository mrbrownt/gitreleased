// vue.config.js

module.exports = {
    devServer: {
        proxy: {
            "^/api": {
                target: "http://host.docker.internal:3000",
            },
            "^/auth": {
                target: "http://host.docker.internal:3000",
            },
        },
    },
}
