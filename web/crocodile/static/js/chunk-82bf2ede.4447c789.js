(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-82bf2ede"],{"1b7f":function(e,t,n){"use strict";var s=n("c8d1"),a=n.n(s);a.a},2017:function(e,t,n){"use strict";var s=n("b12d"),a=n.n(s);a.a},"9ed6":function(e,t,n){"use strict";n.r(t);var s=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{directives:[{name:"loading",rawName:"v-loading",value:e.installloading,expression:"installloading"}],staticClass:"login-container",attrs:{"element-loading-text":"正在安装...","element-loading-spinner":"el-icon-loading","element-loading-background":"rgba(0, 0, 0, 0.8)"}},[n("el-form",{ref:"loginForm",staticClass:"login-form",attrs:{model:e.loginForm,rules:e.loginRules,"auto-complete":"on","label-position":"left"}},[n("div",{staticClass:"title-container"},[n("h3",{staticClass:"title"},[e._v("Crocodile 任务调度平台")]),e._v(" "),n("h6",{directives:[{name:"show",rawName:"v-show",value:e.needinstall,expression:"needinstall"}],staticClass:"installtitle"},[e._v("首次运行请先进行安装操作")])]),e._v(" "),n("el-form-item",{attrs:{prop:"username"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"user"}})],1),e._v(" "),n("el-input",{ref:"username",attrs:{placeholder:"Username",name:"username",type:"text",tabindex:"1","auto-complete":"on"},model:{value:e.loginForm.username,callback:function(t){e.$set(e.loginForm,"username",t)},expression:"loginForm.username"}})],1),e._v(" "),n("el-form-item",{attrs:{prop:"password"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"password"}})],1),e._v(" "),n("el-input",{key:e.passwordType,ref:"password",attrs:{type:e.passwordType,placeholder:"Password",name:"password",tabindex:"2","auto-complete":"on"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleLogin(t)}},model:{value:e.loginForm.password,callback:function(t){e.$set(e.loginForm,"password",t)},expression:"loginForm.password"}})],1),e._v(" "),n("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.needinstall,expression:"needinstall"}]},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"password"}})],1),e._v(" "),n("el-input",{key:e.passwordType2,ref:"password",attrs:{type:e.passwordType2,placeholder:"Password",name:"password",tabindex:"2","auto-complete":"on"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleLogin(t)}},model:{value:e.password2,callback:function(t){e.password2=t},expression:"password2"}})],1),e._v(" "),e.needinstall?n("el-button",{staticStyle:{width:"100%","margin-bottom":"30px"},attrs:{loading:e.loading,type:"primary"},nativeOn:{click:function(t){return t.preventDefault(),e.startinstallcrocodile(t)}}},[e._v("开始安装")]):n("el-button",{staticStyle:{width:"100%","margin-bottom":"30px"},attrs:{loading:e.loading,type:"primary"},nativeOn:{click:function(t){return t.preventDefault(),e.handleLogin(t)}}},[e._v("登陆")]),e._v(" "),n("br")],1)],1)},a=[],o=(n("61f7"),n("b775"));function r(){return Object(o["a"])({url:"/api/v1/install/status",method:"get"})}function i(e){return Object(o["a"])({url:"/api/v1/install",method:"post",data:e})}var l=n("5c96"),d=(n("c24f"),{name:"Login",data:function(){return{loginForm:{username:"",password:""},password2:"",loginRules:{username:[{required:!0,trigger:"blur",message:"请输入用户名"}],password:[{required:!0,trigger:"blur",message:"请输入密码"}]},loading:!1,passwordType:"password",passwordType2:"password",redirect:void 0,needinstall:!1,installloading:!1}},watch:{$route:{handler:function(e){this.redirect=e.query&&e.query.redirect},immediate:!0}},created:function(){this.startqueryinstallstatus()},methods:{startqueryinstallstatus:function(){var e=this;r().then((function(t){10700===t.code&&(e.needinstall=!0,l["Message"].warning(t.msg))}))},startinstallcrocodile:function(){var e=this;if(this.loginForm.password===this.password2){try{window.btoa("".concat(this.loginForm.username,":").concat(this.loginForm.password))}catch(t){return void l["Message"].error("用户名和密码只能使用字母、数字、符号")}this.installloading=!0,i(this.loginForm).then((function(t){0===t.code?(e.startqueryinstallstatus(),e.installloading=!1,e.needinstall=!1,l["Message"].success("恭喜你已经安装成功🎉")):(l["Message"].error(t.msg),e.needinstall=!1)}))}else l["Message"].error("两次密码输入不相同")},handleLogin:function(){var e=this;this.$store.dispatch("user/login",this.loginForm).then((function(){e.$router.push({path:e.redirect||"/"}),e.loading=!1})).catch((function(){e.loading=!1}))}}}),c=d,u=(n("2017"),n("1b7f"),n("2877")),p=Object(u["a"])(c,s,a,!1,null,"d66fbe58",null);t["default"]=p.exports},b12d:function(e,t,n){},c8d1:function(e,t,n){}}]);