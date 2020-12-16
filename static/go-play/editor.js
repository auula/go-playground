// editor核心功能代码
// about: SDing <deen.job@qq.com>

//根据DOM元素的id构造出一个编辑器
const editor = CodeMirror.fromTextArea(document.getElementById("code"), {
    mode: "text/x-go", //实现go代码高亮
    indentUnit: 4, // 缩进单位为4
    lineNumbers: true, //显示行号
    theme: "dracula", //设置主题
    styleActiveLine: true, //当前行高亮
    lineWrapping: true, //代码折叠
    foldGutter: true,
    gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter", ""],
    matchBrackets: true, //括号匹配
    //readOnly: true,        //只读
    // autofocus: true,
    extraKeys: {
        "Ctrl-Space": "autocomplete",
        "Ctrl-T": function () {
            console.log(editor.getValue())
            alert(editor.getValue());
        }
    }, // Tab代码提示
});

/*判断是否处于深色模式*/
if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
    editor.setOption("theme", "dracula");
}
/*判断是否处于浅色模式*/
if (window.matchMedia('(prefers-color-scheme: light)').matches) {
    editor.setOption("theme", "mdn-like");
}
editor.setSize('100%', '100%'); //设置代码框的长宽


// 测试代码内容
const content =
    `package search

// Ctrl-Space autocomplete
// RemoveRep is Deduplication []string
// https://www.fenghong.tech/blog/algorithm/go-slice-deduplicate/
func RemoveRep(a []string) []string {
	if len(a) == 0 {
		return a
	}
	res := make([]string, 0, len(a))
	tmp := map[string]struct{}{}
	for _, v := range a {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}`;

editor.setValue(content); //先代码框的值清空
//editor.setValue(obj.scriptCode);    //给代码框赋值
//editor.setOption("readOnly", true);
//editor.setCode('clear')


/*编辑器上锁*/
function lockEditor() {
    //editor.setOption('readOnly',true)
    if (editor.getOption('readOnly')) {
        editor.setOption('readOnly', false)
        message.left("编辑器解锁.")
    } else {
        editor.setOption('readOnly', true)
        message.left("编辑器上锁.")
    }
}


/*悬浮球*/
const fab = new mdui.Fab('#fab');

function shareCode() {
    mdui.prompt('你的代码分享链接是:',

        function (value) {
            mdui.alert('你输入了：' + value + '，点击了确认按钮');
        },
        function (value) {
            mdui.alert('你输入了：' + value + '，点击了取消按钮');
        },
        {
            type: 'textarea',
            defaultValue: 'https://github.com/higker/go-playground'
        }
    );
}


/*模式切换听器*/
const listeners = {
    dark: function (mediaQueryList) {
        if (mediaQueryList.matches) {
            editor.setOption("theme", "dracula");
        }
    },
    light: function (mediaQueryList) {
        if (mediaQueryList.matches) {
            editor.setOption("theme", "mdn-like");
        }
    }
};

window.matchMedia('(prefers-color-scheme: dark)').addListener(listeners.dark)
window.matchMedia('(prefers-color-scheme: light)').addListener(listeners.light)

