

const $ = mdui.$;
const cs = new mdui.Dialog('#console');
const drawer = new mdui.Drawer('#drawer');


document.getElementById('toggle').addEventListener('click', function () {
    drawer.toggle();
});

const message = {
    right: function (msg) {
        mdui.snackbar({
            message: msg,
            position: 'right-bottom'
        })
    },
    left: function (msg) {
        mdui.snackbar({
            message: msg,
            position: 'left-bottom'
        })
    },
};


console.log(`
░░░░░░░░░░░░░░░░░░░░░░░░▄░░
░░░░░░░░░▐█░░░░░░░░░░░▄▀▒▌░
░░░░░░░░▐▀▒█░░░░░░░░▄▀▒▒▒▐
░░░░░░░▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
░░░░░▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
░░░▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
░░▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒
░░▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐
░▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄
░▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒
▀▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒
单身狗就这样默默地看着你，一句话也不说。
https://github.com/higker/go-playground
`)
