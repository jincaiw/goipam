import{An as e,B as t,Cn as n,Dn as r,Fn as i,Gn as a,Hn as o,L as s,R as c,Rn as l,Tn as u,Un as d,er as f,et as p,fn as m,gn as h,hn as g,ir as _,kn as v,ln as y,nt as b,on as x,rt as S,sn as C,tt as w,z as ee}from"./vendor-DZa9ZXK9.js";import{$ as T,A as E,At as D,C as O,Et as k,F as te,H as ne,J as A,M as j,O as M,Ot as N,S as re,T as ie,Tt as P,V as ae,X as F,at as I,b as oe,bt as se,c as L,et as R,ft as ce,gt as z,ht as B,j as le,jt as V,k as H,kt as U,lt as ue,pt as W,st as G,tt as de,u as K,ut as q,v as fe,vt as pe,w as J,x as me,xt as he,y as ge}from"./vendor-naive-controls-BZmlyIhg.js";import{C as _e,G as ve,S as ye,T as be,W as Y,Z as xe,_t as X,at as Z,it as Se,k as Ce,mt as we,nt as Te,pt as Q,u as $}from"./vendor-naive-core-DREF2OC_.js";import{d as Ee}from"./vendor-naive-data-DhopLyVK.js";var De=Se(`.v-x-scroll`,{overflow:`auto`,scrollbarWidth:`none`},[Se(`&::-webkit-scrollbar`,{width:0,height:0})]),Oe=n({name:`XScroll`,props:{disabled:Boolean,onScroll:Function},setup(){let e=f(null);function t(e){!(e.currentTarget.offsetWidth<e.currentTarget.scrollWidth)||e.deltaY===0||(e.currentTarget.scrollLeft+=e.deltaY+e.deltaX,e.preventDefault())}let n=G();return De.mount({id:`vueuc/x-scroll`,head:!0,anchorMetaName:Z,ssr:n}),Object.assign({selfRef:e,handleWheel:t},{scrollTo(...t){var n;(n=e.value)==null||n.scrollTo(...t)}})},render(){return u(`div`,{ref:`selfRef`,onScroll:this.onScroll,onWheel:this.disabled?void 0:this.handleWheel,class:`v-x-scroll`},this.$slots)}}),ke={name:`zh-CN`,global:{undo:`撤销`,redo:`重做`,confirm:`确认`,clear:`清除`},Popconfirm:{positiveText:`确认`,negativeText:`取消`},Cascader:{placeholder:`请选择`,loading:`加载中`,loadingRequiredMessage:e=>`加载全部 ${e} 的子节点后才可选中`},Time:{dateFormat:`yyyy-MM-dd`,dateTimeFormat:`yyyy-MM-dd HH:mm:ss`},DatePicker:{yearFormat:`yyyy年`,monthFormat:`MMM`,dayFormat:`eeeeee`,yearTypeFormat:`yyyy`,monthTypeFormat:`yyyy-MM`,dateFormat:`yyyy-MM-dd`,dateTimeFormat:`yyyy-MM-dd HH:mm:ss`,quarterFormat:`yyyy-qqq`,weekFormat:`YYYY-w周`,clear:`清除`,now:`此刻`,confirm:`确认`,selectTime:`选择时间`,selectDate:`选择日期`,datePlaceholder:`选择日期`,datetimePlaceholder:`选择日期时间`,monthPlaceholder:`选择月份`,yearPlaceholder:`选择年份`,quarterPlaceholder:`选择季度`,weekPlaceholder:`选择周`,startDatePlaceholder:`开始日期`,endDatePlaceholder:`结束日期`,startDatetimePlaceholder:`开始日期时间`,endDatetimePlaceholder:`结束日期时间`,startMonthPlaceholder:`开始月份`,endMonthPlaceholder:`结束月份`,monthBeforeYear:!1,firstDayOfWeek:0,today:`今天`},DataTable:{checkTableAll:`选择全部表格数据`,uncheckTableAll:`取消选择全部表格数据`,confirm:`确认`,clear:`重置`},LegacyTransfer:{sourceTitle:`源项`,targetTitle:`目标项`},Transfer:{selectAll:`全选`,clearAll:`清除`,unselectAll:`取消全选`,total:e=>`共 ${e} 项`,selected:e=>`已选 ${e} 项`},Empty:{description:`无数据`},Select:{placeholder:`请选择`},TimePicker:{placeholder:`请选择时间`,positiveText:`确认`,negativeText:`取消`,now:`此刻`,clear:`清除`},Pagination:{goto:`跳至`,selectionSuffix:`页`},DynamicTags:{add:`添加`},Log:{loading:`加载中`},Input:{placeholder:`请输入`},InputNumber:{placeholder:`请输入`},DynamicInput:{create:`添加`},ThemeEditor:{title:`主题编辑器`,clearAllVars:`清除全部变量`,clearSearch:`清除搜索`,filterCompName:`过滤组件名`,filterVarName:`过滤变量名`,import:`导入`,export:`导出`,restore:`恢复默认`},Image:{tipPrevious:`上一张（←）`,tipNext:`下一张（→）`,tipCounterclockwise:`向左旋转`,tipClockwise:`向右旋转`,tipZoomOut:`缩小`,tipZoomIn:`放大`,tipDownload:`下载`,tipClose:`关闭（Esc）`,tipOriginalSize:`缩放到原始尺寸`},Heatmap:{less:`少`,more:`多`,monthFormat:`MMM`,weekdayFormat:`eeeeee`}},Ae={lessThanXSeconds:{one:`不到 1 秒`,other:`不到 {{count}} 秒`},xSeconds:{one:`1 秒`,other:`{{count}} 秒`},halfAMinute:`半分钟`,lessThanXMinutes:{one:`不到 1 分钟`,other:`不到 {{count}} 分钟`},xMinutes:{one:`1 分钟`,other:`{{count}} 分钟`},xHours:{one:`1 小时`,other:`{{count}} 小时`},aboutXHours:{one:`大约 1 小时`,other:`大约 {{count}} 小时`},xDays:{one:`1 天`,other:`{{count}} 天`},aboutXWeeks:{one:`大约 1 个星期`,other:`大约 {{count}} 个星期`},xWeeks:{one:`1 个星期`,other:`{{count}} 个星期`},aboutXMonths:{one:`大约 1 个月`,other:`大约 {{count}} 个月`},xMonths:{one:`1 个月`,other:`{{count}} 个月`},aboutXYears:{one:`大约 1 年`,other:`大约 {{count}} 年`},xYears:{one:`1 年`,other:`{{count}} 年`},overXYears:{one:`超过 1 年`,other:`超过 {{count}} 年`},almostXYears:{one:`将近 1 年`,other:`将近 {{count}} 年`}},je=(e,t,n)=>{let r,i=Ae[e];return r=typeof i==`string`?i:t===1?i.one:i.other.replace(`{{count}}`,String(t)),n?.addSuffix?n.comparison&&n.comparison>0?r+`内`:r+`前`:r},Me={date:S({formats:{full:`y'年'M'月'd'日' EEEE`,long:`y'年'M'月'd'日'`,medium:`yyyy-MM-dd`,short:`yy-MM-dd`},defaultWidth:`full`}),time:S({formats:{full:`zzzz a h:mm:ss`,long:`z a h:mm:ss`,medium:`a h:mm:ss`,short:`a h:mm`},defaultWidth:`full`}),dateTime:S({formats:{full:`{{date}} {{time}}`,long:`{{date}} {{time}}`,medium:`{{date}} {{time}}`,short:`{{date}} {{time}}`},defaultWidth:`full`})};function Ne(e,t,n){return Ee(e,t,n)?`eeee p`:e.getTime()>t.getTime()?`'下个'eeee p`:`'上个'eeee p`}var Pe={lastWeek:Ne,yesterday:`'昨天' p`,today:`'今天' p`,tomorrow:`'明天' p`,nextWeek:Ne,other:`PP p`},Fe={name:`zh-CN`,locale:{code:`zh-CN`,formatDistance:je,formatLong:Me,formatRelative:(e,t,n,r)=>{let i=Pe[e];return typeof i==`function`?i(t,n,r):i},localize:{ordinalNumber:(e,t)=>{let n=Number(e);switch(t?.unit){case`date`:return n.toString()+`日`;case`hour`:return n.toString()+`时`;case`minute`:return n.toString()+`分`;case`second`:return n.toString()+`秒`;default:return`第 `+n.toString()}},era:b({values:{narrow:[`前`,`公元`],abbreviated:[`前`,`公元`],wide:[`公元前`,`公元`]},defaultWidth:`wide`}),quarter:b({values:{narrow:[`1`,`2`,`3`,`4`],abbreviated:[`第一季`,`第二季`,`第三季`,`第四季`],wide:[`第一季度`,`第二季度`,`第三季度`,`第四季度`]},defaultWidth:`wide`,argumentCallback:e=>e-1}),month:b({values:{narrow:[`一`,`二`,`三`,`四`,`五`,`六`,`七`,`八`,`九`,`十`,`十一`,`十二`],abbreviated:[`1月`,`2月`,`3月`,`4月`,`5月`,`6月`,`7月`,`8月`,`9月`,`10月`,`11月`,`12月`],wide:[`一月`,`二月`,`三月`,`四月`,`五月`,`六月`,`七月`,`八月`,`九月`,`十月`,`十一月`,`十二月`]},defaultWidth:`wide`}),day:b({values:{narrow:[`日`,`一`,`二`,`三`,`四`,`五`,`六`],short:[`日`,`一`,`二`,`三`,`四`,`五`,`六`],abbreviated:[`周日`,`周一`,`周二`,`周三`,`周四`,`周五`,`周六`],wide:[`星期日`,`星期一`,`星期二`,`星期三`,`星期四`,`星期五`,`星期六`]},defaultWidth:`wide`}),dayPeriod:b({values:{narrow:{am:`上`,pm:`下`,midnight:`凌晨`,noon:`午`,morning:`早`,afternoon:`下午`,evening:`晚`,night:`夜`},abbreviated:{am:`上午`,pm:`下午`,midnight:`凌晨`,noon:`中午`,morning:`早晨`,afternoon:`中午`,evening:`晚上`,night:`夜间`},wide:{am:`上午`,pm:`下午`,midnight:`凌晨`,noon:`中午`,morning:`早晨`,afternoon:`中午`,evening:`晚上`,night:`夜间`}},defaultWidth:`wide`,formattingValues:{narrow:{am:`上`,pm:`下`,midnight:`凌晨`,noon:`午`,morning:`早`,afternoon:`下午`,evening:`晚`,night:`夜`},abbreviated:{am:`上午`,pm:`下午`,midnight:`凌晨`,noon:`中午`,morning:`早晨`,afternoon:`中午`,evening:`晚上`,night:`夜间`},wide:{am:`上午`,pm:`下午`,midnight:`凌晨`,noon:`中午`,morning:`早晨`,afternoon:`中午`,evening:`晚上`,night:`夜间`}},defaultFormattingWidth:`wide`})},match:{ordinalNumber:p({matchPattern:/^(第\s*)?\d+(日|时|分|秒)?/i,parsePattern:/\d+/i,valueCallback:e=>parseInt(e,10)}),era:w({matchPatterns:{narrow:/^(前)/i,abbreviated:/^(前)/i,wide:/^(公元前|公元)/i},defaultMatchWidth:`wide`,parsePatterns:{any:[/^(前)/i,/^(公元)/i]},defaultParseWidth:`any`}),quarter:w({matchPatterns:{narrow:/^[1234]/i,abbreviated:/^第[一二三四]刻/i,wide:/^第[一二三四]刻钟/i},defaultMatchWidth:`wide`,parsePatterns:{any:[/(1|一)/i,/(2|二)/i,/(3|三)/i,/(4|四)/i]},defaultParseWidth:`any`,valueCallback:e=>e+1}),month:w({matchPatterns:{narrow:/^(一|二|三|四|五|六|七|八|九|十[二一]?)/i,abbreviated:/^(一|二|三|四|五|六|七|八|九|十[二一]?|\d|1[0-2])月/i,wide:/^(一|二|三|四|五|六|七|八|九|十[二一]?)月/i},defaultMatchWidth:`wide`,parsePatterns:{narrow:[/^一/i,/^二/i,/^三/i,/^四/i,/^五/i,/^六/i,/^七/i,/^八/i,/^九/i,/^十(?!(一|二))/i,/^十一/i,/^十二/i],any:[/^(一|1(?!\d))/i,/^(二|2)/i,/^(三|3)/i,/^(四|4)/i,/^(五|5)/i,/^(六|6)/i,/^(七|7)/i,/^(八|8)/i,/^(九|9)/i,/^(十(?!(一|二))|10)/i,/^(十一|11)/i,/^(十二|12)/i]},defaultParseWidth:`any`}),day:w({matchPatterns:{narrow:/^[一二三四五六日]/i,short:/^[一二三四五六日]/i,abbreviated:/^周[一二三四五六日]/i,wide:/^星期[一二三四五六日]/i},defaultMatchWidth:`wide`,parsePatterns:{any:[/日/i,/一/i,/二/i,/三/i,/四/i,/五/i,/六/i]},defaultParseWidth:`any`}),dayPeriod:w({matchPatterns:{any:/^(上午?|下午?|午夜|[中正]午|早上?|下午|晚上?|凌晨|)/i},defaultMatchWidth:`any`,parsePatterns:{any:{am:/^上午?/i,pm:/^下午?/i,midnight:/^午夜/i,noon:/^[中正]午/i,morning:/^早上/i,afternoon:/^下午/i,evening:/^晚上?/i,night:/^凌晨/i}},defaultParseWidth:`any`})},options:{weekStartsOn:1,firstWeekContainsDate:4}}};function Ie(e){let{opacityDisabled:t,heightTiny:n,heightSmall:r,heightMedium:i,heightLarge:a,heightHuge:o,primaryColor:s,fontSize:c}=e;return{fontSize:c,textColor:s,sizeTiny:n,sizeSmall:r,sizeMedium:i,sizeLarge:a,sizeHuge:o,color:s,opacitySpinning:t}}var Le={name:`Spin`,common:oe,self:Ie},Re={tabFontSizeSmall:`14px`,tabFontSizeMedium:`14px`,tabFontSizeLarge:`16px`,tabGapSmallLine:`36px`,tabGapMediumLine:`36px`,tabGapLargeLine:`36px`,tabGapSmallLineVertical:`8px`,tabGapMediumLineVertical:`8px`,tabGapLargeLineVertical:`8px`,tabPaddingSmallLine:`6px 0`,tabPaddingMediumLine:`10px 0`,tabPaddingLargeLine:`14px 0`,tabPaddingVerticalSmallLine:`6px 12px`,tabPaddingVerticalMediumLine:`8px 16px`,tabPaddingVerticalLargeLine:`10px 20px`,tabGapSmallBar:`36px`,tabGapMediumBar:`36px`,tabGapLargeBar:`36px`,tabGapSmallBarVertical:`8px`,tabGapMediumBarVertical:`8px`,tabGapLargeBarVertical:`8px`,tabPaddingSmallBar:`4px 0`,tabPaddingMediumBar:`6px 0`,tabPaddingLargeBar:`10px 0`,tabPaddingVerticalSmallBar:`6px 12px`,tabPaddingVerticalMediumBar:`8px 16px`,tabPaddingVerticalLargeBar:`10px 20px`,tabGapSmallCard:`4px`,tabGapMediumCard:`4px`,tabGapLargeCard:`4px`,tabGapSmallCardVertical:`4px`,tabGapMediumCardVertical:`4px`,tabGapLargeCardVertical:`4px`,tabPaddingSmallCard:`8px 16px`,tabPaddingMediumCard:`10px 20px`,tabPaddingLargeCard:`12px 24px`,tabPaddingSmallSegment:`4px 0`,tabPaddingMediumSegment:`6px 0`,tabPaddingLargeSegment:`8px 0`,tabPaddingVerticalLargeSegment:`0 8px`,tabPaddingVerticalSmallCard:`8px 12px`,tabPaddingVerticalMediumCard:`10px 16px`,tabPaddingVerticalLargeCard:`12px 20px`,tabPaddingVerticalSmallSegment:`0 4px`,tabPaddingVerticalMediumSegment:`0 6px`,tabGapSmallSegment:`0`,tabGapMediumSegment:`0`,tabGapLargeSegment:`0`,tabGapSmallSegmentVertical:`0`,tabGapMediumSegmentVertical:`0`,tabGapLargeSegmentVertical:`0`,panePaddingSmall:`8px 0 0 0`,panePaddingMedium:`12px 0 0 0`,panePaddingLarge:`16px 0 0 0`,closeSize:`18px`,closeIconSize:`14px`};function ze(e){let{textColor2:t,primaryColor:n,textColorDisabled:r,closeIconColor:i,closeIconColorHover:a,closeIconColorPressed:o,closeColorHover:s,closeColorPressed:c,tabColor:l,baseColor:u,dividerColor:d,fontWeight:f,textColor1:p,borderRadius:m,fontSize:h,fontWeightStrong:g}=e;return Object.assign(Object.assign({},Re),{colorSegment:l,tabFontSizeCard:h,tabTextColorLine:p,tabTextColorActiveLine:n,tabTextColorHoverLine:n,tabTextColorDisabledLine:r,tabTextColorSegment:p,tabTextColorActiveSegment:t,tabTextColorHoverSegment:t,tabTextColorDisabledSegment:r,tabTextColorBar:p,tabTextColorActiveBar:n,tabTextColorHoverBar:n,tabTextColorDisabledBar:r,tabTextColorCard:p,tabTextColorHoverCard:p,tabTextColorActiveCard:n,tabTextColorDisabledCard:r,barColor:n,closeIconColor:i,closeIconColorHover:a,closeIconColorPressed:o,closeColorHover:s,closeColorPressed:c,closeBorderRadius:m,tabColor:l,tabColorSegment:u,tabBorderColor:d,tabFontWeightActive:f,tabFontWeight:f,tabBorderRadius:m,paneTextColor:t,fontWeightStrong:g})}var Be={name:`Tabs`,common:oe,self:ze};function Ve(e){let{borderRadiusSmall:t,dividerColor:n,hoverColor:r,pressedColor:i,primaryColor:a,textColor3:o,textColor2:s,textColorDisabled:c,fontSize:l}=e;return{fontSize:l,lineHeight:`1.5`,nodeHeight:`30px`,nodeWrapperPadding:`3px 0`,nodeBorderRadius:t,nodeColorHover:r,nodeColorPressed:i,nodeColorActive:z(a,{alpha:.1}),arrowColor:o,nodeTextColor:s,nodeTextColorDisabled:c,loadingColor:a,dropMarkColor:a,lineColor:n}}var He=le({name:`Tree`,common:oe,peers:{Checkbox:K,Scrollbar:ge,Empty:_e},self:Ve}),Ue=P([P(`@keyframes spin-rotate`,`
 from {
 transform: rotate(0);
 }
 to {
 transform: rotate(360deg);
 }
 `),k(`spin-container`,`
 position: relative;
 `,[k(`spin-body`,`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[me()])]),k(`spin-body`,`
 display: inline-flex;
 align-items: center;
 justify-content: center;
 flex-direction: column;
 `),k(`spin`,`
 display: inline-flex;
 height: var(--n-size);
 width: var(--n-size);
 font-size: var(--n-size);
 color: var(--n-color);
 `,[U(`rotate`,`
 animation: spin-rotate 2s linear infinite;
 `)]),k(`spin-description`,`
 display: inline-block;
 font-size: var(--n-font-size);
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 margin-top: 8px;
 `),k(`spin-content`,`
 opacity: 1;
 transition: opacity .3s var(--n-bezier);
 pointer-events: all;
 `,[U(`spinning`,`
 user-select: none;
 -webkit-user-select: none;
 pointer-events: none;
 opacity: var(--n-opacity-spinning);
 `)])]),We={small:20,medium:18,large:16},Ge=n({name:`Spin`,props:Object.assign(Object.assign(Object.assign({},j.props),{contentClass:String,contentStyle:[Object,String],description:String,size:{type:[String,Number],default:`medium`},show:{type:Boolean,default:!0},rotate:{type:Boolean,default:!0},spinning:{type:Boolean,validator:()=>!0,default:void 0},delay:Number}),O),slots:Object,setup(e){let{mergedClsPrefixRef:t,inlineThemeDisabled:n}=ne(e),r=j(`Spin`,`-spin`,Ue,Le,e,t),i=h(()=>{let{size:t}=e,{common:{cubicBezierEaseInOut:n},self:i}=r.value,{opacitySpinning:a,color:o,textColor:s}=i;return{"--n-bezier":n,"--n-opacity-spinning":a,"--n-size":typeof t==`number`?he(t):i[V(`size`,t)],"--n-color":o,"--n-text-color":s}}),a=n?ae(`spin`,h(()=>{let{size:t}=e;return typeof t==`number`?String(t):t[0]}),i,e):void 0,o=Q(e,[`spinning`,`show`]),s=f(!1);return d(t=>{let n;if(o.value){let{delay:r}=e;if(r){n=window.setTimeout(()=>{s.value=!0},r),t(()=>{clearTimeout(n)});return}}s.value=o.value}),{mergedClsPrefix:t,active:s,mergedStrokeWidth:h(()=>{let{strokeWidth:t}=e;if(t!==void 0)return t;let{size:n}=e;return We[typeof n==`number`?`medium`:n]}),cssVars:n?void 0:i,themeClass:a?.themeClass,onRender:a?.onRender}},render(){var e;let{$slots:t,mergedClsPrefix:n,description:r}=this,i=t.icon&&this.rotate,a=(r||t.description)&&u(`div`,{class:`${n}-spin-description`},r||t.description?.call(t)),o=t.icon?u(`div`,{class:[`${n}-spin-body`,this.themeClass]},u(`div`,{class:[`${n}-spin`,i&&`${n}-spin--rotate`],style:t.default?``:this.cssVars},t.icon()),a):u(`div`,{class:[`${n}-spin-body`,this.themeClass]},u(re,{clsPrefix:n,style:t.default?``:this.cssVars,stroke:this.stroke,"stroke-width":this.mergedStrokeWidth,radius:this.radius,scale:this.scale,class:`${n}-spin`}),a);return(e=this.onRender)==null||e.call(this),t.default?u(`div`,{class:[`${n}-spin-container`,this.themeClass],style:this.cssVars},u(`div`,{class:[`${n}-spin-content`,this.active&&`${n}-spin-content--spinning`,this.contentClass],style:this.contentStyle},t),u(x,{name:`fade-in-transition`},{default:()=>this.active?o:null})):o}}),Ke=q(`n-tabs`),qe={tab:[String,Number,Object,Function],name:{type:[String,Number],required:!0},disabled:Boolean,displayDirective:{type:String,default:`if`},closable:{type:Boolean,default:void 0},tabProps:Object,label:[String,Number,Object,Function]},Je=n({__TAB_PANE__:!0,name:`TabPane`,alias:[`TabPanel`],props:qe,slots:Object,setup(e){let t=r(Ke,null);return t||de(`tab-pane`,"`n-tab-pane` must be placed inside `n-tabs`."),{style:t.paneStyleRef,class:t.paneClassRef,mergedClsPrefix:t.mergedClsPrefixRef}},render(){return u(`div`,{class:[`${this.mergedClsPrefix}-tab-pane`,this.class],style:this.style},this.$slots)}}),Ye=n({__TAB__:!0,inheritAttrs:!1,name:`Tab`,props:Object.assign({internalLeftPadded:Boolean,internalAddable:Boolean,internalCreatedByPane:Boolean},ve(qe,[`displayDirective`])),setup(e){let{mergedClsPrefixRef:t,valueRef:n,typeRef:i,closableRef:a,tabStyleRef:o,addTabStyleRef:s,tabClassRef:c,addTabClassRef:l,tabChangeIdRef:u,onBeforeLeaveRef:d,triggerRef:f,handleAdd:p,activateTab:m,handleClose:g}=r(Ke);return{trigger:f,mergedClosable:h(()=>{if(e.internalAddable)return!1;let{closable:t}=e;return t===void 0?a.value:t}),style:o,addStyle:s,tabClass:c,addTabClass:l,clsPrefix:t,value:n,type:i,handleClose(t){t.stopPropagation(),!e.disabled&&g(e.name)},activateTab(){if(e.disabled)return;if(e.internalAddable){p();return}let{name:t}=e,r=++u.id;if(t!==n.value){let{value:i}=d;i?Promise.resolve(i(e.name,n.value)).then(e=>{e&&u.id===r&&m(t)}):m(t)}}}},render(){let{internalAddable:e,clsPrefix:t,name:n,disabled:r,label:i,tab:a,value:o,mergedClosable:s,trigger:c,$slots:{default:l}}=this,d=i??a;return u(`div`,{class:`${t}-tabs-tab-wrapper`},this.internalLeftPadded?u(`div`,{class:`${t}-tabs-tab-pad`}):null,u(`div`,Object.assign({key:n,"data-name":n,"data-disabled":r?!0:void 0},v({class:[`${t}-tabs-tab`,o===n&&`${t}-tabs-tab--active`,r&&`${t}-tabs-tab--disabled`,s&&`${t}-tabs-tab--closable`,e&&`${t}-tabs-tab--addable`,e?this.addTabClass:this.tabClass],onClick:c===`click`?this.activateTab:void 0,onMouseenter:c===`hover`?this.activateTab:void 0,style:e?this.addStyle:this.style},this.internalCreatedByPane?this.tabProps||{}:this.$attrs)),u(`span`,{class:`${t}-tabs-tab__label`},e?u(m,null,u(`div`,{class:`${t}-tabs-tab__height-placeholder`},`\xA0`),u(E,{clsPrefix:t},{default:()=>u(M,null)})):l?l():typeof d==`object`?d:Y(d??n)),s&&this.type===`card`?u(be,{clsPrefix:t,class:`${t}-tabs-tab__close`,onClick:this.handleClose,disabled:r}):null))}}),Xe=k(`tabs`,`
 box-sizing: border-box;
 width: 100%;
 display: flex;
 flex-direction: column;
 transition:
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
`,[U(`segment-type`,[k(`tabs-rail`,[P(`&.transition-disabled`,[k(`tabs-capsule`,`
 transition: none;
 `)])])]),U(`top`,[k(`tab-pane`,`
 padding: var(--n-pane-padding-top) var(--n-pane-padding-right) var(--n-pane-padding-bottom) var(--n-pane-padding-left);
 `)]),U(`left`,[k(`tab-pane`,`
 padding: var(--n-pane-padding-right) var(--n-pane-padding-bottom) var(--n-pane-padding-left) var(--n-pane-padding-top);
 `)]),U(`left, right`,`
 flex-direction: row;
 `,[k(`tabs-bar`,`
 width: 2px;
 right: 0;
 transition:
 top .2s var(--n-bezier),
 max-height .2s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `),k(`tabs-tab`,`
 padding: var(--n-tab-padding-vertical); 
 `)]),U(`right`,`
 flex-direction: row-reverse;
 `,[k(`tab-pane`,`
 padding: var(--n-pane-padding-left) var(--n-pane-padding-top) var(--n-pane-padding-right) var(--n-pane-padding-bottom);
 `),k(`tabs-bar`,`
 left: 0;
 `)]),U(`bottom`,`
 flex-direction: column-reverse;
 justify-content: flex-end;
 `,[k(`tab-pane`,`
 padding: var(--n-pane-padding-bottom) var(--n-pane-padding-right) var(--n-pane-padding-top) var(--n-pane-padding-left);
 `),k(`tabs-bar`,`
 top: 0;
 `)]),k(`tabs-rail`,`
 position: relative;
 padding: 3px;
 border-radius: var(--n-tab-border-radius);
 width: 100%;
 background-color: var(--n-color-segment);
 transition: background-color .3s var(--n-bezier);
 display: flex;
 align-items: center;
 `,[k(`tabs-capsule`,`
 border-radius: var(--n-tab-border-radius);
 position: absolute;
 pointer-events: none;
 background-color: var(--n-tab-color-segment);
 box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .08);
 transition: transform 0.3s var(--n-bezier);
 `),k(`tabs-tab-wrapper`,`
 flex-basis: 0;
 flex-grow: 1;
 display: flex;
 align-items: center;
 justify-content: center;
 `,[k(`tabs-tab`,`
 overflow: hidden;
 border-radius: var(--n-tab-border-radius);
 width: 100%;
 display: flex;
 align-items: center;
 justify-content: center;
 `,[U(`active`,`
 font-weight: var(--n-font-weight-strong);
 color: var(--n-tab-text-color-active);
 `),P(`&:hover`,`
 color: var(--n-tab-text-color-hover);
 `)])])]),U(`flex`,[k(`tabs-nav`,`
 width: 100%;
 position: relative;
 `,[k(`tabs-wrapper`,`
 width: 100%;
 `,[k(`tabs-tab`,`
 margin-right: 0;
 `)])])]),k(`tabs-nav`,`
 box-sizing: border-box;
 line-height: 1.5;
 display: flex;
 transition: border-color .3s var(--n-bezier);
 `,[N(`prefix, suffix`,`
 display: flex;
 align-items: center;
 `),N(`prefix`,`padding-right: 16px;`),N(`suffix`,`padding-left: 16px;`)]),U(`top, bottom`,[P(`>`,[k(`tabs-nav`,[k(`tabs-nav-scroll-wrapper`,[P(`&::before`,`
 top: 0;
 bottom: 0;
 left: 0;
 width: 20px;
 `),P(`&::after`,`
 top: 0;
 bottom: 0;
 right: 0;
 width: 20px;
 `),U(`shadow-start`,[P(`&::before`,`
 box-shadow: inset 10px 0 8px -8px rgba(0, 0, 0, .12);
 `)]),U(`shadow-end`,[P(`&::after`,`
 box-shadow: inset -10px 0 8px -8px rgba(0, 0, 0, .12);
 `)])])])])]),U(`left, right`,[k(`tabs-nav-scroll-content`,`
 flex-direction: column;
 `),P(`>`,[k(`tabs-nav`,[k(`tabs-nav-scroll-wrapper`,[P(`&::before`,`
 top: 0;
 left: 0;
 right: 0;
 height: 20px;
 `),P(`&::after`,`
 bottom: 0;
 left: 0;
 right: 0;
 height: 20px;
 `),U(`shadow-start`,[P(`&::before`,`
 box-shadow: inset 0 10px 8px -8px rgba(0, 0, 0, .12);
 `)]),U(`shadow-end`,[P(`&::after`,`
 box-shadow: inset 0 -10px 8px -8px rgba(0, 0, 0, .12);
 `)])])])])]),k(`tabs-nav-scroll-wrapper`,`
 flex: 1;
 position: relative;
 overflow: hidden;
 `,[k(`tabs-nav-y-scroll`,`
 height: 100%;
 width: 100%;
 overflow-y: auto; 
 scrollbar-width: none;
 `,[P(`&::-webkit-scrollbar, &::-webkit-scrollbar-track-piece, &::-webkit-scrollbar-thumb`,`
 width: 0;
 height: 0;
 display: none;
 `)]),P(`&::before, &::after`,`
 transition: box-shadow .3s var(--n-bezier);
 pointer-events: none;
 content: "";
 position: absolute;
 z-index: 1;
 `)]),k(`tabs-nav-scroll-content`,`
 display: flex;
 position: relative;
 min-width: 100%;
 min-height: 100%;
 width: fit-content;
 box-sizing: border-box;
 `),k(`tabs-wrapper`,`
 display: inline-flex;
 flex-wrap: nowrap;
 position: relative;
 `),k(`tabs-tab-wrapper`,`
 display: flex;
 flex-wrap: nowrap;
 flex-shrink: 0;
 flex-grow: 0;
 `),k(`tabs-tab`,`
 cursor: pointer;
 white-space: nowrap;
 flex-wrap: nowrap;
 display: inline-flex;
 align-items: center;
 color: var(--n-tab-text-color);
 font-size: var(--n-tab-font-size);
 background-clip: padding-box;
 padding: var(--n-tab-padding);
 transition:
 box-shadow .3s var(--n-bezier),
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[U(`disabled`,{cursor:`not-allowed`}),N(`close`,`
 margin-left: 6px;
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `),N(`label`,`
 display: flex;
 align-items: center;
 z-index: 1;
 `)]),k(`tabs-bar`,`
 position: absolute;
 bottom: 0;
 height: 2px;
 border-radius: 1px;
 background-color: var(--n-bar-color);
 transition:
 left .2s var(--n-bezier),
 max-width .2s var(--n-bezier),
 opacity .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `,[P(`&.transition-disabled`,`
 transition: none;
 `),U(`disabled`,`
 background-color: var(--n-tab-text-color-disabled)
 `)]),k(`tabs-pane-wrapper`,`
 position: relative;
 overflow: hidden;
 transition: max-height .2s var(--n-bezier);
 `),k(`tab-pane`,`
 color: var(--n-pane-text-color);
 width: 100%;
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 opacity .2s var(--n-bezier);
 left: 0;
 right: 0;
 top: 0;
 `,[P(`&.next-transition-leave-active, &.prev-transition-leave-active, &.next-transition-enter-active, &.prev-transition-enter-active`,`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 transform .2s var(--n-bezier),
 opacity .2s var(--n-bezier);
 `),P(`&.next-transition-leave-active, &.prev-transition-leave-active`,`
 position: absolute;
 `),P(`&.next-transition-enter-from, &.prev-transition-leave-to`,`
 transform: translateX(32px);
 opacity: 0;
 `),P(`&.next-transition-leave-to, &.prev-transition-enter-from`,`
 transform: translateX(-32px);
 opacity: 0;
 `),P(`&.next-transition-leave-from, &.next-transition-enter-to, &.prev-transition-leave-from, &.prev-transition-enter-to`,`
 transform: translateX(0);
 opacity: 1;
 `)]),k(`tabs-tab-pad`,`
 box-sizing: border-box;
 width: var(--n-tab-gap);
 flex-grow: 0;
 flex-shrink: 0;
 `),U(`line-type, bar-type`,[k(`tabs-tab`,`
 font-weight: var(--n-tab-font-weight);
 box-sizing: border-box;
 vertical-align: bottom;
 `,[P(`&:hover`,{color:`var(--n-tab-text-color-hover)`}),U(`active`,`
 color: var(--n-tab-text-color-active);
 font-weight: var(--n-tab-font-weight-active);
 `),U(`disabled`,{color:`var(--n-tab-text-color-disabled)`})])]),k(`tabs-nav`,[U(`line-type`,[U(`top`,[N(`prefix, suffix`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),k(`tabs-nav-scroll-content`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),k(`tabs-bar`,`
 bottom: -1px;
 `)]),U(`left`,[N(`prefix, suffix`,`
 border-right: 1px solid var(--n-tab-border-color);
 `),k(`tabs-nav-scroll-content`,`
 border-right: 1px solid var(--n-tab-border-color);
 `),k(`tabs-bar`,`
 right: -1px;
 `)]),U(`right`,[N(`prefix, suffix`,`
 border-left: 1px solid var(--n-tab-border-color);
 `),k(`tabs-nav-scroll-content`,`
 border-left: 1px solid var(--n-tab-border-color);
 `),k(`tabs-bar`,`
 left: -1px;
 `)]),U(`bottom`,[N(`prefix, suffix`,`
 border-top: 1px solid var(--n-tab-border-color);
 `),k(`tabs-nav-scroll-content`,`
 border-top: 1px solid var(--n-tab-border-color);
 `),k(`tabs-bar`,`
 top: -1px;
 `)]),N(`prefix, suffix`,`
 transition: border-color .3s var(--n-bezier);
 `),k(`tabs-nav-scroll-content`,`
 transition: border-color .3s var(--n-bezier);
 `),k(`tabs-bar`,`
 border-radius: 0;
 `)]),U(`card-type`,[N(`prefix, suffix`,`
 transition: border-color .3s var(--n-bezier);
 `),k(`tabs-pad`,`
 flex-grow: 1;
 transition: border-color .3s var(--n-bezier);
 `),k(`tabs-tab-pad`,`
 transition: border-color .3s var(--n-bezier);
 `),k(`tabs-tab`,`
 font-weight: var(--n-tab-font-weight);
 border: 1px solid var(--n-tab-border-color);
 background-color: var(--n-tab-color);
 box-sizing: border-box;
 position: relative;
 vertical-align: bottom;
 display: flex;
 justify-content: space-between;
 font-size: var(--n-tab-font-size);
 color: var(--n-tab-text-color);
 `,[U(`addable`,`
 padding-left: 8px;
 padding-right: 8px;
 font-size: 16px;
 justify-content: center;
 `,[N(`height-placeholder`,`
 width: 0;
 font-size: var(--n-tab-font-size);
 `),D(`disabled`,[P(`&:hover`,`
 color: var(--n-tab-text-color-hover);
 `)])]),U(`closable`,`padding-right: 8px;`),U(`active`,`
 background-color: #0000;
 font-weight: var(--n-tab-font-weight-active);
 color: var(--n-tab-text-color-active);
 `),U(`disabled`,`color: var(--n-tab-text-color-disabled);`)])]),U(`left, right`,`
 flex-direction: column; 
 `,[N(`prefix, suffix`,`
 padding: var(--n-tab-padding-vertical);
 `),k(`tabs-wrapper`,`
 flex-direction: column;
 `),k(`tabs-tab-wrapper`,`
 flex-direction: column;
 `,[k(`tabs-tab-pad`,`
 height: var(--n-tab-gap-vertical);
 width: 100%;
 `)])]),U(`top`,[U(`card-type`,[k(`tabs-scroll-padding`,`border-bottom: 1px solid var(--n-tab-border-color);`),N(`prefix, suffix`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),k(`tabs-tab`,`
 border-top-left-radius: var(--n-tab-border-radius);
 border-top-right-radius: var(--n-tab-border-radius);
 `,[U(`active`,`
 border-bottom: 1px solid #0000;
 `)]),k(`tabs-tab-pad`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),k(`tabs-pad`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `)])]),U(`left`,[U(`card-type`,[k(`tabs-scroll-padding`,`border-right: 1px solid var(--n-tab-border-color);`),N(`prefix, suffix`,`
 border-right: 1px solid var(--n-tab-border-color);
 `),k(`tabs-tab`,`
 border-top-left-radius: var(--n-tab-border-radius);
 border-bottom-left-radius: var(--n-tab-border-radius);
 `,[U(`active`,`
 border-right: 1px solid #0000;
 `)]),k(`tabs-tab-pad`,`
 border-right: 1px solid var(--n-tab-border-color);
 `),k(`tabs-pad`,`
 border-right: 1px solid var(--n-tab-border-color);
 `)])]),U(`right`,[U(`card-type`,[k(`tabs-scroll-padding`,`border-left: 1px solid var(--n-tab-border-color);`),N(`prefix, suffix`,`
 border-left: 1px solid var(--n-tab-border-color);
 `),k(`tabs-tab`,`
 border-top-right-radius: var(--n-tab-border-radius);
 border-bottom-right-radius: var(--n-tab-border-radius);
 `,[U(`active`,`
 border-left: 1px solid #0000;
 `)]),k(`tabs-tab-pad`,`
 border-left: 1px solid var(--n-tab-border-color);
 `),k(`tabs-pad`,`
 border-left: 1px solid var(--n-tab-border-color);
 `)])]),U(`bottom`,[U(`card-type`,[k(`tabs-scroll-padding`,`border-top: 1px solid var(--n-tab-border-color);`),N(`prefix, suffix`,`
 border-top: 1px solid var(--n-tab-border-color);
 `),k(`tabs-tab`,`
 border-bottom-left-radius: var(--n-tab-border-radius);
 border-bottom-right-radius: var(--n-tab-border-radius);
 `,[U(`active`,`
 border-top: 1px solid #0000;
 `)]),k(`tabs-tab-pad`,`
 border-top: 1px solid var(--n-tab-border-color);
 `),k(`tabs-pad`,`
 border-top: 1px solid var(--n-tab-border-color);
 `)])])])]),Ze=t,Qe=n({name:`Tabs`,props:Object.assign(Object.assign({},j.props),{value:[String,Number],defaultValue:[String,Number],trigger:{type:String,default:`click`},type:{type:String,default:`bar`},closable:Boolean,justifyContent:String,size:String,placement:{type:String,default:`top`},tabStyle:[String,Object],tabClass:String,addTabStyle:[String,Object],addTabClass:String,barWidth:Number,paneClass:String,paneStyle:[String,Object],paneWrapperClass:String,paneWrapperStyle:[String,Object],addable:[Boolean,Object],tabsPadding:{type:Number,default:0},animated:Boolean,onBeforeLeave:Function,onAdd:Function,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],onClose:[Function,Array],labelSize:String,activeName:[String,Number],onActiveNameChange:[Function,Array]}),slots:Object,setup(t,{slots:n}){let{mergedClsPrefixRef:r,inlineThemeDisabled:a,mergedComponentPropsRef:s}=ne(t),c=j(`Tabs`,`-tabs`,Xe,Be,t,r),u=f(null),p=f(null),m=f(null),g=f(null),v=f(null),y=f(null),b=f(!0),x=f(!0),S=Q(t,[`labelSize`,`size`]),C=h(()=>S.value?S.value:s?.value?.Tabs?.size||`medium`),w=Q(t,[`activeName`,`value`]),ee=f(w.value??t.defaultValue??(n.default?T(n.default())[0]?.props?.name:null)),E=ce(w,ee),D={id:0},O=h(()=>{if(!(!t.justifyContent||t.type===`card`))return{display:`flex`,justifyContent:t.justifyContent}});o(E,()=>{D.id=0,N(),re()});function k(){let{value:e}=E;return e===null?null:u.value?.querySelector(`[data-name="${e}"]`)}function te(e){if(t.type===`card`)return;let{value:n}=p;if(!n)return;let i=n.style.opacity===`0`;if(e){let a=`${r.value}-tabs-bar--disabled`,{barWidth:o,placement:s}=t;if(e.dataset.disabled===`true`?n.classList.add(a):n.classList.remove(a),[`top`,`bottom`].includes(s)){if(M([`top`,`maxHeight`,`height`]),typeof o==`number`&&e.offsetWidth>=o){let t=Math.floor((e.offsetWidth-o)/2)+e.offsetLeft;n.style.left=`${t}px`,n.style.maxWidth=`${o}px`}else n.style.left=`${e.offsetLeft}px`,n.style.maxWidth=`${e.offsetWidth}px`;n.style.width=`8192px`,i&&(n.style.transition=`none`),n.offsetWidth,i&&(n.style.transition=``,n.style.opacity=`1`)}else{if(M([`left`,`maxWidth`,`width`]),typeof o==`number`&&e.offsetHeight>=o){let t=Math.floor((e.offsetHeight-o)/2)+e.offsetTop;n.style.top=`${t}px`,n.style.maxHeight=`${o}px`}else n.style.top=`${e.offsetTop}px`,n.style.maxHeight=`${e.offsetHeight}px`;n.style.height=`8192px`,i&&(n.style.transition=`none`),n.offsetHeight,i&&(n.style.transition=``,n.style.opacity=`1`)}}}function A(){if(t.type===`card`)return;let{value:e}=p;e&&(e.style.opacity=`0`)}function M(e){let{value:t}=p;if(t)for(let n of e)t.style[n]=``}function N(){if(t.type===`card`)return;let e=k();e?te(e):A()}function re(){let e=v.value?.$el;if(!e)return;let t=k();if(!t)return;let{scrollLeft:n,offsetWidth:r}=e,{offsetLeft:i,offsetWidth:a}=t;n>i?e.scrollTo({top:0,left:i,behavior:`smooth`}):i+a>n+r&&e.scrollTo({top:0,left:i+a-r,behavior:`smooth`})}let ie=f(null),P=0,F=null;function I(e){let t=ie.value;if(t){P=e.getBoundingClientRect().height;let n=`${P}px`,r=()=>{t.style.height=n,t.style.maxHeight=n};F?(r(),F(),F=null):F=r}}function oe(e){let t=ie.value;if(t){let n=e.getBoundingClientRect().height,r=()=>{document.body.offsetHeight,t.style.maxHeight=`${n}px`,t.style.height=`${Math.max(P,n)}px`};F?(F(),F=null,r()):F=r}}function L(){let e=ie.value;if(e){e.style.maxHeight=``,e.style.height=``;let{paneWrapperStyle:n}=t;if(typeof n==`string`)e.style.cssText=n;else if(n){let{maxHeight:t,height:r}=n;t!==void 0&&(e.style.maxHeight=t),r!==void 0&&(e.style.height=r)}}}let z={value:[]},B=f(`next`);function le(e){let t=E.value,n=`next`;for(let r of z.value){if(r===t)break;if(r===e){n=`prev`;break}}B.value=n,H(e)}function H(e){let{onActiveNameChange:n,onUpdateValue:r,"onUpdate:value":i}=t;n&&R(n,e),r&&R(r,e),i&&R(i,e),ee.value=e}function U(e){let{onClose:n}=t;n&&R(n,e)}let ue=!0;function W(){let{value:e}=p;if(!e)return;ue||=!1;let t=`transition-disabled`;e.classList.add(t),N(),e.classList.remove(t)}let G=f(null);function de({transitionDisabled:e}){let t=u.value;if(!t)return;e&&t.classList.add(`transition-disabled`);let n=k();n&&G.value&&(G.value.style.width=`${n.offsetWidth}px`,G.value.style.height=`${n.offsetHeight}px`,G.value.style.transform=`translateX(${n.offsetLeft-pe(getComputedStyle(t).paddingLeft)}px)`,e&&G.value.offsetWidth),e&&t.classList.remove(`transition-disabled`)}o([E],()=>{t.type===`segment`&&e(()=>{de({transitionDisabled:!1})})}),i(()=>{t.type===`segment`&&de({transitionDisabled:!0})});let K=0;function q(e){if(e.contentRect.width===0&&e.contentRect.height===0||K===e.contentRect.width)return;K=e.contentRect.width;let{type:n}=t;if((n===`line`||n===`bar`)&&(ue||t.justifyContent?.startsWith(`space`))&&W(),n!==`segment`){let{placement:e}=t;_e((e===`top`||e===`bottom`?v.value?.$el:y.value)||null)}}let fe=Ze(q,64);o([()=>t.justifyContent,()=>t.size],()=>{e(()=>{let{type:e}=t;(e===`line`||e===`bar`)&&W()})});let J=f(!1);function me(e){let{target:n,contentRect:{width:r,height:i}}=e,a=n.parentElement.parentElement.offsetWidth,o=n.parentElement.parentElement.offsetHeight,{placement:s}=t;if(!J.value)s===`top`||s===`bottom`?a<r&&(J.value=!0):o<i&&(J.value=!0);else{let{value:e}=g;if(!e)return;s===`top`||s===`bottom`?a-r>e.$el.offsetWidth&&(J.value=!1):o-i>e.$el.offsetHeight&&(J.value=!1)}_e(v.value?.$el||null)}let he=Ze(me,64);function ge(){let{onAdd:n}=t;n&&n(),e(()=>{let e=k(),{value:t}=v;!e||!t||t.scrollTo({left:e.offsetLeft,top:0,behavior:`smooth`})})}function _e(e){if(!e)return;let{placement:n}=t;if(n===`top`||n===`bottom`){let{scrollLeft:t,scrollWidth:n,offsetWidth:r}=e;b.value=t<=0,x.value=t+r>=n}else{let{scrollTop:t,scrollHeight:n,offsetHeight:r}=e;b.value=t<=0,x.value=t+r>=n}}let ve=Ze(e=>{_e(e.target)},64);l(Ke,{triggerRef:_(t,`trigger`),tabStyleRef:_(t,`tabStyle`),tabClassRef:_(t,`tabClass`),addTabStyleRef:_(t,`addTabStyle`),addTabClassRef:_(t,`addTabClass`),paneClassRef:_(t,`paneClass`),paneStyleRef:_(t,`paneStyle`),mergedClsPrefixRef:r,typeRef:_(t,`type`),closableRef:_(t,`closable`),valueRef:E,tabChangeIdRef:D,onBeforeLeaveRef:_(t,`onBeforeLeave`),activateTab:le,handleClose:U,handleAdd:ge}),we(()=>{N(),re()}),d(()=>{let{value:e}=m;if(!e)return;let{value:t}=r,n=`${t}-tabs-nav-scroll-wrapper--shadow-start`,i=`${t}-tabs-nav-scroll-wrapper--shadow-end`;b.value?e.classList.remove(n):e.classList.add(n),x.value?e.classList.remove(i):e.classList.add(i)});let ye={syncBarPosition:()=>{N()}},be=()=>{de({transitionDisabled:!0})},Y=h(()=>{let{value:e}=C,{type:n}=t,r=`${e}${{card:`Card`,bar:`Bar`,line:`Line`,segment:`Segment`}[n]}`,{self:{barColor:i,closeIconColor:a,closeIconColorHover:o,closeIconColorPressed:s,tabColor:l,tabBorderColor:u,paneTextColor:d,tabFontWeight:f,tabBorderRadius:p,tabFontWeightActive:m,colorSegment:h,fontWeightStrong:g,tabColorSegment:_,closeSize:v,closeIconSize:y,closeColorHover:b,closeColorPressed:x,closeBorderRadius:S,[V(`panePadding`,e)]:w,[V(`tabPadding`,r)]:ee,[V(`tabPaddingVertical`,r)]:T,[V(`tabGap`,r)]:E,[V(`tabGap`,`${r}Vertical`)]:D,[V(`tabTextColor`,n)]:O,[V(`tabTextColorActive`,n)]:k,[V(`tabTextColorHover`,n)]:te,[V(`tabTextColorDisabled`,n)]:ne,[V(`tabFontSize`,e)]:A},common:{cubicBezierEaseInOut:j}}=c.value;return{"--n-bezier":j,"--n-color-segment":h,"--n-bar-color":i,"--n-tab-font-size":A,"--n-tab-text-color":O,"--n-tab-text-color-active":k,"--n-tab-text-color-disabled":ne,"--n-tab-text-color-hover":te,"--n-pane-text-color":d,"--n-tab-border-color":u,"--n-tab-border-radius":p,"--n-close-size":v,"--n-close-icon-size":y,"--n-close-color-hover":b,"--n-close-color-pressed":x,"--n-close-border-radius":S,"--n-close-icon-color":a,"--n-close-icon-color-hover":o,"--n-close-icon-color-pressed":s,"--n-tab-color":l,"--n-tab-font-weight":f,"--n-tab-font-weight-active":m,"--n-tab-padding":ee,"--n-tab-padding-vertical":T,"--n-tab-gap":E,"--n-tab-gap-vertical":D,"--n-pane-padding-left":se(w,`left`),"--n-pane-padding-right":se(w,`right`),"--n-pane-padding-top":se(w,`top`),"--n-pane-padding-bottom":se(w,`bottom`),"--n-font-weight-strong":g,"--n-tab-color-segment":_}}),xe=a?ae(`tabs`,h(()=>`${C.value[0]}${t.type[0]}`),Y,t):void 0;return Object.assign({mergedClsPrefix:r,mergedValue:E,renderedNames:new Set,segmentCapsuleElRef:G,tabsPaneWrapperRef:ie,tabsElRef:u,barElRef:p,addTabInstRef:g,xScrollInstRef:v,scrollWrapperElRef:m,addTabFixed:J,tabWrapperStyle:O,handleNavResize:fe,mergedSize:C,handleScroll:ve,handleTabsResize:he,cssVars:a?void 0:Y,themeClass:xe?.themeClass,animationDirection:B,renderNameListRef:z,yScrollElRef:y,handleSegmentResize:be,onAnimationBeforeLeave:I,onAnimationEnter:oe,onAnimationAfterEnter:L,onRender:xe?.onRender},ye)},render(){let{mergedClsPrefix:e,type:t,placement:n,addTabFixed:r,addable:i,mergedSize:a,renderNameListRef:o,onRender:s,paneWrapperClass:c,paneWrapperStyle:l,$slots:{default:d,prefix:f,suffix:p}}=this;s?.();let m=d?T(d()).filter(e=>e.type.__TAB_PANE__===!0):[],h=d?T(d()).filter(e=>e.type.__TAB__===!0):[],g=!h.length,_=t===`card`,v=t===`segment`,y=!_&&!v&&this.justifyContent;o.value=[];let b=()=>{let t=u(`div`,{style:this.tabWrapperStyle,class:`${e}-tabs-wrapper`},y?null:u(`div`,{class:`${e}-tabs-scroll-padding`,style:n===`top`||n===`bottom`?{width:`${this.tabsPadding}px`}:{height:`${this.tabsPadding}px`}}),g?m.map((e,t)=>(o.value.push(e.props.name),nt(u(Ye,Object.assign({},e.props,{internalCreatedByPane:!0,internalLeftPadded:t!==0&&(!y||y===`center`||y===`start`||y===`end`)}),e.children?{default:e.children.tab}:void 0)))):h.map((e,t)=>(o.value.push(e.props.name),nt(t!==0&&!y?tt(e):e))),!r&&i&&_?et(i,(g?m.length:h.length)!==0):null,y?null:u(`div`,{class:`${e}-tabs-scroll-padding`,style:{width:`${this.tabsPadding}px`}}));return u(`div`,{ref:`tabsElRef`,class:`${e}-tabs-nav-scroll-content`},_&&i?u(I,{onResize:this.handleTabsResize},{default:()=>t}):t,_?u(`div`,{class:`${e}-tabs-pad`}):null,_?null:u(`div`,{ref:`barElRef`,class:`${e}-tabs-bar`}))},x=v?`top`:n;return u(`div`,{class:[`${e}-tabs`,this.themeClass,`${e}-tabs--${t}-type`,`${e}-tabs--${a}-size`,y&&`${e}-tabs--flex`,`${e}-tabs--${x}`],style:this.cssVars},u(`div`,{class:[`${e}-tabs-nav--${t}-type`,`${e}-tabs-nav--${x}`,`${e}-tabs-nav`]},F(f,t=>t&&u(`div`,{class:`${e}-tabs-nav__prefix`},t)),v?u(I,{onResize:this.handleSegmentResize},{default:()=>u(`div`,{class:`${e}-tabs-rail`,ref:`tabsElRef`},u(`div`,{class:`${e}-tabs-capsule`,ref:`segmentCapsuleElRef`},u(`div`,{class:`${e}-tabs-wrapper`},u(`div`,{class:`${e}-tabs-tab`}))),g?m.map((e,t)=>(o.value.push(e.props.name),u(Ye,Object.assign({},e.props,{internalCreatedByPane:!0,internalLeftPadded:t!==0}),e.children?{default:e.children.tab}:void 0))):h.map((e,t)=>(o.value.push(e.props.name),t===0?e:tt(e))))}):u(I,{onResize:this.handleNavResize},{default:()=>u(`div`,{class:`${e}-tabs-nav-scroll-wrapper`,ref:`scrollWrapperElRef`},[`top`,`bottom`].includes(x)?u(Oe,{ref:`xScrollInstRef`,onScroll:this.handleScroll},{default:b}):u(`div`,{class:`${e}-tabs-nav-y-scroll`,onScroll:this.handleScroll,ref:`yScrollElRef`},b()))}),r&&i&&_?et(i,!0):null,F(p,t=>t&&u(`div`,{class:`${e}-tabs-nav__suffix`},t))),g&&(this.animated&&(x===`top`||x===`bottom`)?u(`div`,{ref:`tabsPaneWrapperRef`,style:l,class:[`${e}-tabs-pane-wrapper`,c]},$e(m,this.mergedValue,this.renderedNames,this.onAnimationBeforeLeave,this.onAnimationEnter,this.onAnimationAfterEnter,this.animationDirection)):$e(m,this.mergedValue,this.renderedNames)))}});function $e(e,t,n,r,i,o,s){let c=[];return e.forEach(e=>{let{name:r,displayDirective:i,"display-directive":o}=e.props,s=e=>i===e||o===e,l=t===r;if(e.key!==void 0&&(e.key=r),l||s(`show`)||s(`show:lazy`)&&n.has(r)){n.has(r)||n.add(r);let t=!s(`if`);c.push(t?a(e,[[y,l]]):e)}}),s?u(C,{name:`${s}-transition`,onBeforeLeave:r,onEnter:i,onAfterEnter:o},{default:()=>c}):c}function et(e,t){return u(Ye,{ref:`addTabInstRef`,key:`__addable`,name:`__addable`,internalCreatedByPane:!0,internalAddable:!0,internalLeftPadded:t,disabled:typeof e==`object`&&e.disabled})}function tt(e){let t=g(e);return t.props?t.props.internalLeftPadded=!0:t.props={internalLeftPadded:!0},t}function nt(e){return Array.isArray(e.dynamicProps)?e.dynamicProps.includes(`internalLeftPadded`)||e.dynamicProps.push(`internalLeftPadded`):e.dynamicProps=[`internalLeftPadded`],e}var rt=q(`n-tree-select`);function it({position:e,offsetLevel:t,indent:n,el:r}){let i={position:`absolute`,boxSizing:`border-box`,right:0};if(e===`inside`)i.left=0,i.top=0,i.bottom=0,i.borderRadius=`inherit`,i.boxShadow=`inset 0 0 0 2px var(--n-drop-mark-color)`;else{let a=e===`before`?`top`:`bottom`;i[a]=0,i.left=`${r.offsetLeft+6-t*n}px`,i.height=`2px`,i.backgroundColor=`var(--n-drop-mark-color)`,i.transformOrigin=a,i.borderRadius=`1px`,i.transform=e===`before`?`translateY(-4px)`:`translateY(4px)`}return u(`div`,{style:i})}function at({dropPosition:e,node:t}){return t.isLeaf===!1||t.children?!0:e!==`inside`}var ot=q(`n-tree`);function st({props:e,fNodesRef:t,mergedExpandedKeysRef:n,mergedSelectedKeysRef:i,mergedCheckedKeysRef:a,handleCheck:o,handleSelect:s,handleSwitcherClick:c}){let{value:l}=i,u=r(rt,null),d=u?u.pendingNodeKeyRef:f(l.length?l[l.length-1]:null);function p(r){if(!e.keyboard)return{enterBehavior:null};let{value:i}=d,l=null;if(i===null){if((r.key===`ArrowDown`||r.key===`ArrowUp`)&&r.preventDefault(),[`ArrowDown`,`ArrowUp`,`ArrowLeft`,`ArrowRight`].includes(r.key)&&i===null){let{value:e}=t,n=0;for(;n<e.length;){if(!e[n].disabled){d.value=e[n].key;break}n+=1}}}else{let{value:u}=t,f=u.findIndex(e=>e.key===i);if(!~f)return{enterBehavior:null};if(r.key===`Enter`){let t=u[f];switch(l=e.overrideDefaultNodeClickBehavior?.call(e,{option:t.rawNode})||null,l){case`toggleCheck`:o(t,!a.value.includes(t.key));break;case`toggleSelect`:s(t);break;case`toggleExpand`:c(t);break;case`none`:break;default:l=`default`,s(t)}}else if(r.key===`ArrowDown`)for(r.preventDefault(),f+=1;f<u.length;){if(!u[f].disabled){d.value=u[f].key;break}f+=1}else if(r.key===`ArrowUp`)for(r.preventDefault(),--f;f>=0;){if(!u[f].disabled){d.value=u[f].key;break}--f}else if(r.key===`ArrowLeft`){let e=u[f];if(e.isLeaf||!n.value.includes(i)){let t=e.getParent();t&&(d.value=t.key)}else c(e)}else if(r.key===`ArrowRight`){let e=u[f];if(e.isLeaf)return{enterBehavior:null};if(!n.value.includes(i))c(e);else for(f+=1;f<u.length;){if(!u[f].disabled){d.value=u[f].key;break}f+=1}}}return{enterBehavior:l}}return{pendingNodeKeyRef:d,handleKeydown:p}}var ct=n({name:`NTreeNodeCheckbox`,props:{clsPrefix:{type:String,required:!0},indent:{type:Number,required:!0},right:Boolean,focusable:Boolean,disabled:Boolean,checked:Boolean,indeterminate:Boolean,onCheck:Function},setup(e){let t=r(ot);function n(t){let{onCheck:n}=e;n&&n(t)}function i(e){n(e)}return{handleUpdateValue:i,mergedTheme:t.mergedThemeRef}},render(){let{clsPrefix:e,mergedTheme:t,checked:n,indeterminate:r,disabled:i,focusable:a,indent:o,handleUpdateValue:s}=this;return u(`span`,{class:[`${e}-tree-node-checkbox`,this.right&&`${e}-tree-node-checkbox--right`],style:{width:`${o}px`},"data-checkbox":!0},u(L,{focusable:a,disabled:i,theme:t.peers.Checkbox,themeOverrides:t.peerOverrides.Checkbox,checked:n,indeterminate:r,onUpdateChecked:s}))}}),lt=n({name:`TreeNodeContent`,props:{clsPrefix:{type:String,required:!0},disabled:Boolean,checked:Boolean,selected:Boolean,onClick:Function,onDragstart:Function,tmNode:{type:Object,required:!0},nodeProps:Object},setup(e){let{renderLabelRef:t,renderPrefixRef:n,renderSuffixRef:i,labelFieldRef:a}=r(ot),o=f(null);function s(t){let{onClick:n}=e;n&&n(t)}function c(e){s(e)}return{selfRef:o,renderLabel:t,renderPrefix:n,renderSuffix:i,labelField:a,handleClick:c}},render(){let{clsPrefix:e,labelField:t,nodeProps:n,checked:r=!1,selected:i=!1,renderLabel:a,renderPrefix:o,renderSuffix:s,handleClick:c,onDragstart:l,tmNode:{rawNode:d,rawNode:{prefix:f,suffix:p,[t]:m}}}=this;return u(`span`,Object.assign({},n,{ref:`selfRef`,class:[`${e}-tree-node-content`,n?.class],onClick:c,draggable:l===void 0?void 0:!0,onDragstart:l}),o||f?u(`div`,{class:`${e}-tree-node-content__prefix`},o?o({option:d,selected:i,checked:r}):Y(f)):null,u(`div`,{class:`${e}-tree-node-content__text`},a?a({option:d,selected:i,checked:r}):Y(m)),s||p?u(`div`,{class:`${e}-tree-node-content__suffix`},s?s({option:d,selected:i,checked:r}):Y(p)):null)}}),ut=n({name:`NTreeSwitcher`,props:{clsPrefix:{type:String,required:!0},indent:{type:Number,required:!0},expanded:Boolean,selected:Boolean,hide:Boolean,loading:Boolean,onClick:Function,tmNode:{type:Object,required:!0}},setup(e){let{renderSwitcherIconRef:t,spinPropsRef:n}=r(ot,null);return()=>{let{clsPrefix:r,expanded:i,hide:a,indent:o,onClick:s}=e;return u(`span`,{"data-switcher":!0,class:[`${r}-tree-node-switcher`,i&&`${r}-tree-node-switcher--expanded`,a&&`${r}-tree-node-switcher--hide`],style:{width:`${o}px`},onClick:s},u(`div`,{class:`${r}-tree-node-switcher__icon`},u(H,null,{default:()=>{if(e.loading)return u(re,Object.assign({clsPrefix:r,key:`loading`,radius:85,strokeWidth:20},n?.value));let{value:i}=t;return i?i({expanded:e.expanded,selected:e.selected,option:e.tmNode.rawNode}):u(E,{clsPrefix:r,key:`switcher`},{default:()=>u(Ce,null)})}})))}}});function dt(e){return h(()=>e.leafOnly?`child`:e.checkStrategy)}function ft(e,t){return!!e.rawNode[t]}function pt(e,t,n,r){e?.forEach(e=>{n(e),pt(e[t],t,n,r),r(e)})}function mt(e,t,n,r,i){let a=new Set,o=new Set,s=[];return pt(e,r,e=>{if(s.push(e),i(t,e)){o.add(e[n]);for(let e=s.length-2;e>=0;--e)if(!a.has(s[e][n]))a.add(s[e][n]);else return}},()=>{s.pop()}),{expandedKeys:Array.from(a),highlightKeySet:o}}if(ue&&Image){let e=new Image;e.src=`data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==`}function ht(e,t,n,r,i){let a=new Set,o=new Set,s=new Set,c=[],l=[],u=[];function d(e){e.forEach(e=>{if(u.push(e),t(n,e)){a.add(e[r]),s.add(e[r]);for(let e=u.length-2;e>=0;--e){let t=u[e][r];if(!o.has(t))o.add(t),a.has(t)&&a.delete(t);else break}}let c=e[i];c&&d(c),u.pop()})}d(e);function f(e,t){e.forEach(e=>{let n=e[r],s=a.has(n),l=o.has(n);if(!s&&!l)return;let u=e[i];if(u)if(s)t.push(e);else{c.push(n);let r=Object.assign(Object.assign({},e),{[i]:[]});t.push(r),f(u,r[i])}else t.push(e)})}return f(e,l),{filteredTree:l,highlightKeySet:s,expandedKeys:c}}var gt=n({name:`TreeNode`,props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(e){let t=r(ot),{droppingNodeParentRef:n,droppingMouseNodeRef:a,draggingNodeRef:o,droppingPositionRef:s,droppingOffsetLevelRef:c,nodePropsRef:l,indentRef:d,blockLineRef:p,checkboxPlacementRef:m,checkOnClickRef:g,disabledFieldRef:_,showLineRef:v,renderSwitcherIconRef:y,overrideDefaultNodeClickBehaviorRef:b}=t,x=W(()=>!!e.tmNode.rawNode.checkboxDisabled),S=W(()=>ft(e.tmNode,_.value)),C=W(()=>t.disabledRef.value||S.value),w=h(()=>{let{value:t}=l;if(t)return t({option:e.tmNode.rawNode})}),ee=f(null),T={value:null};i(()=>{T.value=ee.value.$el});function E(){let n=()=>{let{tmNode:n}=e;if(!n.isLeaf&&!n.shallowLoaded){if(!t.loadingKeysRef.value.has(n.key))t.loadingKeysRef.value.add(n.key);else return;let{onLoadRef:{value:e}}=t;e&&e(n.rawNode).then(e=>{e!==!1&&t.handleSwitcherClick(n)}).finally(()=>{t.loadingKeysRef.value.delete(n.key)})}else t.handleSwitcherClick(n)};y.value?setTimeout(n,0):n()}let D=W(()=>!S.value&&t.selectableRef.value&&(t.internalTreeSelect?t.mergedCheckStrategyRef.value!==`child`||t.multipleRef.value&&t.cascadeRef.value||e.tmNode.isLeaf:!0)),O=W(()=>t.checkableRef.value&&(t.cascadeRef.value||t.mergedCheckStrategyRef.value!==`child`||e.tmNode.isLeaf)),k=W(()=>t.displayedCheckedKeysRef.value.includes(e.tmNode.key)),te=W(()=>{let{value:t}=O;if(!t)return!1;let{value:n}=g,{tmNode:r}=e;return typeof n==`boolean`?!r.disabled&&n:n(e.tmNode.rawNode)});function ne(n){let{value:r}=t.expandOnClickRef,{value:i}=D,{value:a}=te;if(!i&&!r&&!a||X(n,`checkbox`)||X(n,`switcher`))return;let{tmNode:o}=e;i&&t.handleSelect(o),r&&!o.isLeaf&&E(),a&&N(!k.value)}function A(n){var r,i;if(!(X(n,`checkbox`)||X(n,`switcher`))){if(!C.value){let r=b.value,i=!1;if(r)switch(r({option:e.tmNode.rawNode})){case`toggleCheck`:i=!0,N(!k.value);break;case`toggleSelect`:i=!0,t.handleSelect(e.tmNode);break;case`toggleExpand`:i=!0,E(),i=!0;break;case`none`:i=!0,i=!0;return;default:break}i||ne(n)}(i=(r=w.value)?.onClick)==null||i.call(r,n)}}function j(e){p.value||A(e)}function M(e){p.value&&A(e)}function N(n){t.handleCheck(e.tmNode,n)}function re(n){t.handleDragStart({event:n,node:e.tmNode})}function ie(n){n.currentTarget===n.target&&t.handleDragEnter({event:n,node:e.tmNode})}function P(n){n.preventDefault(),t.handleDragOver({event:n,node:e.tmNode})}function ae(n){t.handleDragEnd({event:n,node:e.tmNode})}function F(n){n.currentTarget===n.target&&t.handleDragLeave({event:n,node:e.tmNode})}function I(n){n.preventDefault(),s.value!==null&&t.handleDrop({event:n,node:e.tmNode,dropPosition:s.value})}let oe=h(()=>{let{clsPrefix:t}=e,{value:n}=d;if(v.value){let r=[],i=e.tmNode.parent;for(;i;)i.isLastChild?r.push(u(`div`,{class:`${t}-tree-node-indent`},u(`div`,{style:{width:`${n}px`}}))):r.push(u(`div`,{class:[`${t}-tree-node-indent`,`${t}-tree-node-indent--show-line`]},u(`div`,{style:{width:`${n}px`}}))),i=i.parent;return r.reverse()}else return B(e.tmNode.level,u(`div`,{class:`${e.clsPrefix}-tree-node-indent`},u(`div`,{style:{width:`${n}px`}})))});return{showDropMark:W(()=>{let{value:t}=o;if(!t)return;let{value:n}=s;if(!n)return;let{value:r}=a;if(!r)return;let{tmNode:i}=e;return i.key===r.key}),showDropMarkAsParent:W(()=>{let{value:t}=n;if(!t)return!1;let{tmNode:r}=e,{value:i}=s;return i===`before`||i===`after`?t.key===r.key:!1}),pending:W(()=>t.pendingNodeKeyRef.value===e.tmNode.key),loading:W(()=>t.loadingKeysRef.value.has(e.tmNode.key)),highlight:W(()=>t.highlightKeySetRef.value?.has(e.tmNode.key)),checked:k,indeterminate:W(()=>t.displayedIndeterminateKeysRef.value.includes(e.tmNode.key)),selected:W(()=>t.mergedSelectedKeysRef.value.includes(e.tmNode.key)),expanded:W(()=>t.mergedExpandedKeysRef.value.includes(e.tmNode.key)),disabled:C,checkable:O,mergedCheckOnClick:te,checkboxDisabled:x,selectable:D,expandOnClick:t.expandOnClickRef,internalScrollable:t.internalScrollableRef,draggable:t.draggableRef,blockLine:p,nodeProps:w,checkboxFocusable:t.internalCheckboxFocusableRef,droppingPosition:s,droppingOffsetLevel:c,indent:d,checkboxPlacement:m,showLine:v,contentInstRef:ee,contentElRef:T,indentNodes:oe,handleCheck:N,handleDrop:I,handleDragStart:re,handleDragEnter:ie,handleDragOver:P,handleDragEnd:ae,handleDragLeave:F,handleLineClick:M,handleContentClick:j,handleSwitcherClick:E}},render(){let{tmNode:e,clsPrefix:t,checkable:n,expandOnClick:r,selectable:i,selected:a,checked:o,highlight:s,draggable:c,blockLine:l,indent:d,indentNodes:f,disabled:p,pending:m,internalScrollable:h,nodeProps:g,checkboxPlacement:_}=this,v=c&&!p?{onDragenter:this.handleDragEnter,onDragleave:this.handleDragLeave,onDragend:this.handleDragEnd,onDrop:this.handleDrop,onDragover:this.handleDragOver}:void 0,y=h?xe(e.key):void 0,b=_===`right`,x=n?u(ct,{indent:d,right:b,focusable:this.checkboxFocusable,disabled:p||this.checkboxDisabled,clsPrefix:t,checked:this.checked,indeterminate:this.indeterminate,onCheck:this.handleCheck}):null;return u(`div`,Object.assign({class:`${t}-tree-node-wrapper`},v),u(`div`,Object.assign({},l?g:void 0,{class:[`${t}-tree-node`,{[`${t}-tree-node--selected`]:a,[`${t}-tree-node--checkable`]:n,[`${t}-tree-node--highlight`]:s,[`${t}-tree-node--pending`]:m,[`${t}-tree-node--disabled`]:p,[`${t}-tree-node--selectable`]:i,[`${t}-tree-node--clickable`]:i||r||this.mergedCheckOnClick},g?.class],"data-key":y,draggable:c&&l,onClick:this.handleLineClick,onDragstart:c&&l&&!p?this.handleDragStart:void 0}),f,e.isLeaf&&this.showLine?u(`div`,{class:[`${t}-tree-node-indent`,`${t}-tree-node-indent--show-line`,e.isLeaf&&`${t}-tree-node-indent--is-leaf`,e.isLastChild&&`${t}-tree-node-indent--last-child`]},u(`div`,{style:{width:`${d}px`}})):u(ut,{clsPrefix:t,expanded:this.expanded,selected:a,loading:this.loading,hide:e.isLeaf,tmNode:this.tmNode,indent:d,onClick:this.handleSwitcherClick}),b?null:x,u(lt,{ref:`contentInstRef`,clsPrefix:t,checked:o,selected:a,onClick:this.handleContentClick,nodeProps:l?void 0:g,onDragstart:c&&!l&&!p?this.handleDragStart:void 0,tmNode:e}),c?this.showDropMark?it({el:this.contentElRef.value,position:this.droppingPosition,offsetLevel:this.droppingOffsetLevel,indent:d}):this.showDropMarkAsParent?it({el:this.contentElRef.value,position:`inside`,offsetLevel:this.droppingOffsetLevel,indent:d}):null:null,b?x:null))}}),_t=n({name:`TreeMotionWrapper`,props:{clsPrefix:{type:String,required:!0},height:Number,nodes:{type:Array,required:!0},mode:{type:String,required:!0},onAfterEnter:{type:Function,required:!0}},render(){let{clsPrefix:e}=this;return u(J,{onAfterEnter:this.onAfterEnter,appear:!0,reverse:this.mode===`collapse`},{default:()=>u(`div`,{class:[`${e}-tree-motion-wrapper`,`${e}-tree-motion-wrapper--${this.mode}`],style:{height:he(this.height)}},this.nodes.map(t=>u(gt,{clsPrefix:e,tmNode:t})))})}}),vt=ie(),yt=k(`tree`,`
 font-size: var(--n-font-size);
 outline: none;
`,[P(`ul, li`,`
 margin: 0;
 padding: 0;
 list-style: none;
 `),P(`>`,[k(`tree-node`,[P(`&:first-child`,`margin-top: 0;`)])]),k(`tree-motion-wrapper`,[U(`expand`,[$({duration:`0.2s`})]),U(`collapse`,[$({duration:`0.2s`,reverse:!0})])]),k(`tree-node-wrapper`,`
 box-sizing: border-box;
 padding: var(--n-node-wrapper-padding);
 `),k(`tree-node`,`
 position: relative;
 display: flex;
 border-radius: var(--n-node-border-radius);
 transition: background-color .3s var(--n-bezier);
 `,[U(`highlight`,[k(`tree-node-content`,[N(`text`,`border-bottom-color: var(--n-node-text-color-disabled);`)])]),U(`disabled`,[k(`tree-node-content`,`
 color: var(--n-node-text-color-disabled);
 cursor: not-allowed;
 `)]),D(`disabled`,[U(`clickable`,[k(`tree-node-content`,`
 cursor: pointer;
 `)])])]),U(`block-node`,[k(`tree-node-content`,`
 flex: 1;
 min-width: 0;
 `)]),D(`block-line`,[k(`tree-node`,[D(`disabled`,[k(`tree-node-content`,[P(`&:hover`,`background: var(--n-node-color-hover);`)]),U(`selectable`,[k(`tree-node-content`,[P(`&:active`,`background: var(--n-node-color-pressed);`)])]),U(`pending`,[k(`tree-node-content`,`
 background: var(--n-node-color-hover);
 `)]),U(`selected`,[k(`tree-node-content`,`background: var(--n-node-color-active);`)])]),U(`selected`,[k(`tree-node-content`,`background: var(--n-node-color-active);`)])])]),U(`block-line`,[k(`tree-node`,[D(`disabled`,[P(`&:hover`,`background: var(--n-node-color-hover);`),U(`pending`,`
 background: var(--n-node-color-hover);
 `),U(`selectable`,[D(`selected`,[P(`&:active`,`background: var(--n-node-color-pressed);`)])]),U(`selected`,`background: var(--n-node-color-active);`)]),U(`selected`,`background: var(--n-node-color-active);`),U(`disabled`,`
 cursor: not-allowed;
 `)])]),U(`ellipsis`,[k(`tree-node`,[k(`tree-node-content`,`
 overflow: hidden;
 `,[N(`text`,`
 text-overflow: ellipsis;
 white-space: nowrap;
 overflow: hidden;
 `)])])]),k(`tree-node-indent`,`
 flex-grow: 0;
 flex-shrink: 0;
 `,[U(`show-line`,`position: relative`,[P(`&::before`,`
 position: absolute;
 left: 50%;
 border-left: 1px solid var(--n-line-color);
 transition: border-color .3s var(--n-bezier);
 transform: translate(-50%);
 content: "";
 top: var(--n-line-offset-top);
 bottom: var(--n-line-offset-bottom);
 `),U(`last-child`,[P(`&::before`,`
 bottom: 50%;
 `)]),U(`is-leaf`,[P(`&::after`,`
 position: absolute;
 content: "";
 left: calc(50% + 0.5px);
 right: 0;
 bottom: 50%;
 transition: border-color .3s var(--n-bezier);
 border-bottom: 1px solid var(--n-line-color);
 `)])]),D(`show-line`,`height: 0;`)]),k(`tree-node-switcher`,`
 cursor: pointer;
 display: inline-flex;
 flex-shrink: 0;
 height: var(--n-node-content-height);
 align-items: center;
 justify-content: center;
 transition: transform .15s var(--n-bezier);
 vertical-align: bottom;
 `,[N(`icon`,`
 position: relative;
 height: 14px;
 width: 14px;
 display: flex;
 color: var(--n-arrow-color);
 transition: color .3s var(--n-bezier);
 font-size: 14px;
 `,[k(`icon`,[vt]),k(`base-loading`,`
 color: var(--n-loading-color);
 position: absolute;
 left: 0;
 top: 0;
 right: 0;
 bottom: 0;
 `,[vt]),k(`base-icon`,[vt])]),U(`hide`,`visibility: hidden;`),U(`expanded`,`transform: rotate(90deg);`)]),k(`tree-node-checkbox`,`
 display: inline-flex;
 height: var(--n-node-content-height);
 vertical-align: bottom;
 align-items: center;
 justify-content: center;
 `),k(`tree-node-content`,`
 user-select: none;
 position: relative;
 display: inline-flex;
 align-items: center;
 min-height: var(--n-node-content-height);
 box-sizing: border-box;
 line-height: var(--n-line-height);
 vertical-align: bottom;
 padding: 0 6px 0 4px;
 cursor: default;
 border-radius: var(--n-node-border-radius);
 color: var(--n-node-text-color);
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[P(`&:last-child`,`margin-bottom: 0;`),N(`prefix`,`
 display: inline-flex;
 margin-right: 8px;
 `),N(`text`,`
 border-bottom: 1px solid #0000;
 transition: border-color .3s var(--n-bezier);
 flex-grow: 1;
 max-width: 100%;
 `),N(`suffix`,`
 display: inline-flex;
 `)]),N(`empty`,`margin: auto;`)]),bt=function(e,t,n,r){function i(e){return e instanceof n?e:new n(function(t){t(e)})}return new(n||=Promise)(function(n,a){function o(e){try{c(r.next(e))}catch(e){a(e)}}function s(e){try{c(r.throw(e))}catch(e){a(e)}}function c(e){e.done?n(e.value):i(e.value).then(o,s)}c((r=r.apply(e,t||[])).next())})};function xt(e,t,n,r){return{getIsGroup(){return!1},getKey(t){return t[e]},getChildren:r||(e=>e[t]),getDisabled(e){return!!(e[n]||e.checkboxDisabled)}}}var St={allowCheckingNotLoaded:Boolean,filter:Function,defaultExpandAll:Boolean,expandedKeys:Array,keyField:{type:String,default:`key`},labelField:{type:String,default:`label`},childrenField:{type:String,default:`children`},disabledField:{type:String,default:`disabled`},defaultExpandedKeys:{type:Array,default:()=>[]},indent:{type:Number,default:24},indeterminateKeys:Array,renderSwitcherIcon:Function,onUpdateIndeterminateKeys:[Function,Array],"onUpdate:indeterminateKeys":[Function,Array],onUpdateExpandedKeys:[Function,Array],"onUpdate:expandedKeys":[Function,Array],overrideDefaultNodeClickBehavior:Function},Ct=n({name:`Tree`,props:Object.assign(Object.assign(Object.assign(Object.assign({},j.props),{accordion:Boolean,showIrrelevantNodes:{type:Boolean,default:!0},data:{type:Array,default:()=>[]},expandOnDragenter:{type:Boolean,default:!0},expandOnClick:Boolean,checkOnClick:{type:[Boolean,Function],default:!1},cancelable:{type:Boolean,default:!0},checkable:Boolean,draggable:Boolean,blockNode:Boolean,blockLine:Boolean,showLine:Boolean,disabled:Boolean,checkedKeys:Array,defaultCheckedKeys:{type:Array,default:()=>[]},selectedKeys:Array,defaultSelectedKeys:{type:Array,default:()=>[]},multiple:Boolean,pattern:{type:String,default:``},onLoad:Function,cascade:Boolean,selectable:{type:Boolean,default:!0},scrollbarProps:Object,allowDrop:{type:Function,default:at},animated:{type:Boolean,default:!0},ellipsis:Boolean,checkboxPlacement:{type:String,default:`left`},virtualScroll:Boolean,watchProps:Array,renderLabel:Function,renderPrefix:Function,renderSuffix:Function,nodeProps:Function,keyboard:{type:Boolean,default:!0},getChildren:Function,onDragenter:[Function,Array],onDragleave:[Function,Array],onDragend:[Function,Array],onDragstart:[Function,Array],onDragover:[Function,Array],onDrop:[Function,Array],onUpdateCheckedKeys:[Function,Array],"onUpdate:checkedKeys":[Function,Array],onUpdateSelectedKeys:[Function,Array],"onUpdate:selectedKeys":[Function,Array]}),St),{internalTreeSelect:Boolean,internalScrollable:Boolean,internalScrollablePadding:String,internalRenderEmpty:Function,internalHighlightKeySet:Object,internalUnifySelectCheck:Boolean,internalCheckboxFocusable:{type:Boolean,default:!0},internalFocusable:{type:Boolean,default:!0},checkStrategy:{type:String,default:`all`},spinProps:Object,leafOnly:Boolean}),slots:Object,setup(t){let{mergedClsPrefixRef:n,inlineThemeDisabled:i,mergedRtlRef:a,mergedComponentPropsRef:u}=ne(t),p=te(`Tree`,a,n),m=j(`Tree`,`-tree`,yt,He,t,n),g=h(()=>u?.value?.Tree?.renderEmpty),v=f(null),y=f(null),b=f(null);function x(){return b.value?.listElRef}function S(){return b.value?.itemsElRef}let C=h(()=>{let{filter:e}=t;if(e)return e;let{labelField:n}=t;return(e,t)=>{if(!e.length)return!0;let r=t[n];return typeof r==`string`?r.toLowerCase().includes(e.toLowerCase()):!1}}),w=h(()=>{let{pattern:e}=t;return!e||!e.length||!C.value?{filteredTree:t.data,highlightKeySet:null,expandedKeys:void 0}:ht(t.data,C.value,e,t.keyField,t.childrenField)}),T=h(()=>s(t.showIrrelevantNodes?t.data:w.value.filteredTree,xt(t.keyField,t.childrenField,t.disabledField,t.getChildren))),E=r(rt,null),D=t.internalTreeSelect?E.dataTreeMate:h(()=>t.showIrrelevantNodes?T.value:s(t.data,xt(t.keyField,t.childrenField,t.disabledField,t.getChildren))),{watchProps:O}=t,k=f([]);O?.includes(`defaultCheckedKeys`)?d(()=>{k.value=t.defaultCheckedKeys}):k.value=t.defaultCheckedKeys;let A=ce(_(t,`checkedKeys`),k),M=h(()=>D.value.getCheckedKeys(A.value,{cascade:t.cascade,allowNotLoaded:t.allowCheckingNotLoaded})),N=dt(t),re=h(()=>M.value.checkedKeys),ie=h(()=>{let{indeterminateKeys:e}=t;return e===void 0?M.value.indeterminateKeys:e}),P=f([]);O?.includes(`defaultSelectedKeys`)?d(()=>{P.value=t.defaultSelectedKeys}):P.value=t.defaultSelectedKeys;let F=ce(_(t,`selectedKeys`),P),I=f([]),oe=e=>{I.value=t.defaultExpandAll?D.value.getNonLeafKeys():e===void 0?t.defaultExpandedKeys:e};O?.includes(`defaultExpandedKeys`)?d(()=>{oe(void 0)}):d(()=>{oe(t.defaultExpandedKeys)});let L=ce(_(t,`expandedKeys`),I),z=h(()=>T.value.getFlattenedNodes(L.value)),{pendingNodeKeyRef:B,handleKeydown:le}=st({props:t,mergedCheckedKeysRef:A,mergedSelectedKeysRef:F,fNodesRef:z,mergedExpandedKeysRef:L,handleCheck:ze,handleSelect:Ue,handleSwitcherClick:Ve}),V=null,H=null,U=f(new Set),ue=ce(h(()=>t.internalHighlightKeySet||w.value.highlightKeySet),U),W=f(new Set),G=h(()=>L.value.filter(e=>!W.value.has(e))),de=0,K=f(null),q=f(null),fe=f(null),J=f(null),me=f(0),ge=h(()=>{let{value:e}=q;return e?e.parent:null}),_e=!1;o(_(t,`data`),()=>{_e=!0,e(()=>{_e=!1}),W.value.clear(),B.value=null,Fe()},{deep:!1});let ve=!1,ye=()=>{ve=!0,e(()=>{ve=!1})},be;o(_(t,`pattern`),(e,n)=>{if(t.showIrrelevantNodes)if(be=void 0,e){let{expandedKeys:e,highlightKeySet:n}=mt(t.data,t.pattern,t.keyField,t.childrenField,C.value);U.value=n,ye(),$(e,Q(e),{node:null,action:`filter`})}else U.value=new Set;else if(!e.length)be!==void 0&&(ye(),$(be,Q(be),{node:null,action:`filter`}));else{n.length||(be=L.value);let{expandedKeys:e}=w.value;e!==void 0&&(ye(),$(e,Q(e),{node:null,action:`filter`}))}});function Y(e){return bt(this,void 0,void 0,function*(){let{onLoad:n}=t;if(!n){yield Promise.resolve();return}let{value:r}=W;if(!r.has(e.key)){r.add(e.key);try{(yield n(e.rawNode))===!1&&Re()}catch(e){console.error(e),Re()}r.delete(e.key)}})}d(()=>{var e;let{value:t}=T;if(!t)return;let{getNode:n}=t;(e=L.value)==null||e.forEach(e=>{let t=n(e);t&&!t.shallowLoaded&&Y(t)})});let X=f(!1),Z=f([]);o(G,(n,r)=>{if(!t.animated||ve){e(we);return}if(_e)return;let i=pe(m.value.self.nodeHeight),a=new Set(r),o=null,s=null;for(let e of n)if(!a.has(e)){if(o!==null)return;o=e}let l=new Set(n);for(let e of r)if(!l.has(e)){if(s!==null)return;s=e}if(o===null&&s===null)return;let{virtualScroll:u}=t,d=(u?b.value.listElRef:v.value).offsetHeight,f=Math.ceil(d/i)+1,p;if(o!==null&&(p=r),s!==null&&(p=p===void 0?n:p.filter(e=>e!==s)),X.value=!0,Z.value=T.value.getFlattenedNodes(p),o!==null){let e=Z.value.findIndex(e=>e.key===o);if(~e){let t=Z.value[e].children;if(t){let r=c(t,n);Z.value.splice(e+1,0,{__motion:!0,mode:`expand`,height:u?r.length*i:void 0,nodes:u?r.slice(0,f):r})}}}if(s!==null){let e=Z.value.findIndex(e=>e.key===s);if(~e){let t=Z.value[e].children;if(!t)return;X.value=!0;let r=c(t,n);Z.value.splice(e+1,0,{__motion:!0,mode:`collapse`,height:u?r.length*i:void 0,nodes:u?r.slice(0,f):r})}}});let Se=h(()=>ee(z.value)),Ce=h(()=>X.value?Z.value:z.value);function we(){let{value:e}=y;e&&e.sync()}function Te(){X.value=!1,t.virtualScroll&&e(we)}function Q(e){let{getNode:t}=D.value;return e.map(e=>t(e)?.rawNode||null)}function $(e,n,r){let{"onUpdate:expandedKeys":i,onUpdateExpandedKeys:a}=t;I.value=e,i&&R(i,e,n,r),a&&R(a,e,n,r)}function Ee(e,n,r){let{"onUpdate:checkedKeys":i,onUpdateCheckedKeys:a}=t;k.value=e,a&&R(a,e,n,r),i&&R(i,e,n,r)}function De(e,n){let{"onUpdate:indeterminateKeys":r,onUpdateIndeterminateKeys:i}=t;r&&R(r,e,n),i&&R(i,e,n)}function Oe(e,n,r){let{"onUpdate:selectedKeys":i,onUpdateSelectedKeys:a}=t;P.value=e,a&&R(a,e,n,r),i&&R(i,e,n,r)}function ke(e){let{onDragenter:n}=t;n&&R(n,e)}function Ae(e){let{onDragleave:n}=t;n&&R(n,e)}function je(e){let{onDragend:n}=t;n&&R(n,e)}function Me(e){let{onDragstart:n}=t;n&&R(n,e)}function Ne(e){let{onDragover:n}=t;n&&R(n,e)}function Pe(e){let{onDrop:n}=t;n&&R(n,e)}function Fe(){Ie(),Le()}function Ie(){K.value=null}function Le(){me.value=0,q.value=null,fe.value=null,J.value=null,Re()}function Re(){V&&=(window.clearTimeout(V),null),H=null}function ze(e,n){if(t.disabled||ft(e,t.disabledField))return;if(t.internalUnifySelectCheck&&!t.multiple){Ue(e);return}let r=n?`check`:`uncheck`,{checkedKeys:i,indeterminateKeys:a}=D.value[r](e.key,re.value,{cascade:t.cascade,checkStrategy:N.value,allowNotLoaded:t.allowCheckingNotLoaded});Ee(i,Q(i),{node:e.rawNode,action:r}),De(a,Q(a))}function Be(e){if(t.disabled)return;let{key:n}=e,{value:r}=L,i=r.findIndex(e=>e===n);if(~i){let t=Array.from(r);t.splice(i,1),$(t,Q(t),{node:e.rawNode,action:`collapse`})}else{let i=T.value.getNode(n);if(!i||i.isLeaf)return;let a;if(t.accordion){let t=new Set(e.siblings.map(({key:e})=>e));a=r.filter(e=>!t.has(e)),a.push(n)}else a=r.concat(n);$(a,Q(a),{node:e.rawNode,action:`expand`})}}function Ve(e){t.disabled||X.value||Be(e)}function Ue(e){if(!(t.disabled||!t.selectable)){if(B.value=e.key,t.internalUnifySelectCheck){let{value:{checkedKeys:n,indeterminateKeys:r}}=M;t.multiple?ze(e,!(n.includes(e.key)||r.includes(e.key))):Ee([e.key],Q([e.key]),{node:e.rawNode,action:`check`})}if(t.multiple){let n=Array.from(F.value),r=n.findIndex(t=>t===e.key);~r?t.cancelable&&n.splice(r,1):~r||n.push(e.key),Oe(n,Q(n),{node:e.rawNode,action:~r?`unselect`:`select`})}else F.value.includes(e.key)?t.cancelable&&Oe([],[],{node:e.rawNode,action:`unselect`}):Oe([e.key],Q([e.key]),{node:e.rawNode,action:`select`})}}function We(e){if(V&&=(window.clearTimeout(V),null),e.isLeaf)return;H=e.key;let t=()=>{if(H!==e.key)return;let{value:t}=fe;if(t&&t.key===e.key&&!L.value.includes(e.key)){let t=L.value.concat(e.key);$(t,Q(t),{node:e.rawNode,action:`expand`})}V=null,H=null};V=e.shallowLoaded?window.setTimeout(()=>{t()},1e3):window.setTimeout(()=>{Y(e).then(()=>{t()})},1e3)}function Ge({event:e,node:n}){!t.draggable||t.disabled||ft(n,t.disabledField)||(Xe({event:e,node:n},!1),ke({event:e,node:n.rawNode}))}function Ke({event:e,node:n}){!t.draggable||t.disabled||ft(n,t.disabledField)||Ae({event:e,node:n.rawNode})}function qe(e){e.target===e.currentTarget&&Le()}function Je({event:e,node:n}){Fe(),!(!t.draggable||t.disabled||ft(n,t.disabledField))&&je({event:e,node:n.rawNode})}function Ye({event:e,node:n}){!t.draggable||t.disabled||ft(n,t.disabledField)||(de=e.clientX,K.value=n,Me({event:e,node:n.rawNode}))}function Xe({event:e,node:n},r=!0){if(!t.draggable||t.disabled||ft(n,t.disabledField))return;let{value:i}=K;if(!i)return;let{allowDrop:a,indent:o}=t;r&&Ne({event:e,node:n.rawNode});let{height:s,top:c}=e.currentTarget.getBoundingClientRect(),l=e.clientY-c,u;u=a({node:n.rawNode,dropPosition:`inside`,phase:`drag`})?l<=8?`before`:l>=s-8?`after`:`inside`:l<=s/2?`before`:`after`;let{value:d}=Se,f,p,m=d(n.key);if(m===null){Le();return}let h=!1;u===`inside`?(f=n,p=`inside`):u===`before`?n.isFirstChild?(f=n,p=`before`):(f=z.value[m-1],p=`after`):(f=n,p=`after`),!f.isLeaf&&L.value.includes(f.key)&&(h=!0,p===`after`&&(f=z.value[m+1],f?p=`before`:(f=n,p=`inside`)));let g=f;if(fe.value=g,!h&&i.isLastChild&&i.key===f.key&&(p=`after`),p===`after`){let t=de-e.clientX,n=0;for(;t>=o/2&&f.parent!==null&&f.isLastChild&&n<1;)t-=o,n+=1,f=f.parent;me.value=n}else me.value=0;if((i.contains(f)||p===`inside`&&i.parent?.key===f.key)&&!(i.key===g.key&&i.key===f.key)){Le();return}if(!a({node:f.rawNode,dropPosition:p,phase:`drag`})){Le();return}if(i.key===f.key)Re();else if(H!==f.key)if(p===`inside`){if(t.expandOnDragenter){if(We(f),!f.shallowLoaded&&H!==f.key){Fe();return}}else if(!f.shallowLoaded){Fe();return}}else Re();else p!==`inside`&&Re();J.value=p,q.value=f}function Ze({event:e,node:n,dropPosition:r}){if(!t.draggable||t.disabled||ft(n,t.disabledField))return;let{value:i}=K,{value:a}=q,{value:o}=J;if(!(!i||!a||!o)&&t.allowDrop({node:a.rawNode,dropPosition:o,phase:`drag`})&&i.key!==a.key){if(o===`before`){let e=i.getNext({includeDisabled:!0});if(e&&e.key===a.key){Le();return}}if(o===`after`){let e=i.getPrev({includeDisabled:!0});if(e&&e.key===a.key){Le();return}}Pe({event:e,node:a.rawNode,dragNode:i.rawNode,dropPosition:r}),Fe()}}function Qe(){we()}function $e(){we()}function et(e){if(t.virtualScroll||t.internalScrollable){let{value:t}=y;if((t?.containerRef)?.contains(e.relatedTarget))return;B.value=null}else{let{value:t}=v;if(t?.contains(e.relatedTarget))return;B.value=null}}o(B,e=>{var n;if(e!==null){if(t.virtualScroll)(n=b.value)==null||n.scrollTo({key:e});else if(t.internalScrollable){let{value:t}=y;if(t===null)return;let n=t.contentRef?.querySelector(`[data-key="${xe(e)}"]`);if(!n)return;t.scrollTo({el:n})}}}),l(ot,{loadingKeysRef:W,highlightKeySetRef:ue,displayedCheckedKeysRef:re,displayedIndeterminateKeysRef:ie,mergedSelectedKeysRef:F,mergedExpandedKeysRef:L,mergedThemeRef:m,mergedCheckStrategyRef:N,nodePropsRef:_(t,`nodeProps`),disabledRef:_(t,`disabled`),checkableRef:_(t,`checkable`),selectableRef:_(t,`selectable`),expandOnClickRef:_(t,`expandOnClick`),onLoadRef:_(t,`onLoad`),draggableRef:_(t,`draggable`),blockLineRef:_(t,`blockLine`),indentRef:_(t,`indent`),cascadeRef:_(t,`cascade`),checkOnClickRef:_(t,`checkOnClick`),checkboxPlacementRef:t.checkboxPlacement,droppingMouseNodeRef:fe,droppingNodeParentRef:ge,draggingNodeRef:K,droppingPositionRef:J,droppingOffsetLevelRef:me,fNodesRef:z,pendingNodeKeyRef:B,showLineRef:_(t,`showLine`),disabledFieldRef:_(t,`disabledField`),internalScrollableRef:_(t,`internalScrollable`),internalCheckboxFocusableRef:_(t,`internalCheckboxFocusable`),internalTreeSelect:t.internalTreeSelect,renderLabelRef:_(t,`renderLabel`),renderPrefixRef:_(t,`renderPrefix`),renderSuffixRef:_(t,`renderSuffix`),renderSwitcherIconRef:_(t,`renderSwitcherIcon`),labelFieldRef:_(t,`labelField`),multipleRef:_(t,`multiple`),overrideDefaultNodeClickBehaviorRef:_(t,`overrideDefaultNodeClickBehavior`),spinPropsRef:_(t,`spinProps`),handleSwitcherClick:Ve,handleDragEnd:Je,handleDragEnter:Ge,handleDragLeave:Ke,handleDragStart:Ye,handleDrop:Ze,handleDragOver:Xe,handleSelect:Ue,handleCheck:ze});function tt(e,t){var n,r;typeof e==`number`?(n=b.value)==null||n.scrollTo(e,t||0):(r=b.value)==null||r.scrollTo(e)}let nt={handleKeydown:le,scrollTo:tt,getCheckedData:()=>{if(!t.checkable)return{keys:[],options:[]};let{checkedKeys:e}=M.value;return{keys:e,options:Q(e)}},getIndeterminateData:()=>{if(!t.checkable)return{keys:[],options:[]};let{indeterminateKeys:e}=M.value;return{keys:e,options:Q(e)}}},it=h(()=>{let{common:{cubicBezierEaseInOut:e},self:{fontSize:t,nodeBorderRadius:n,nodeColorHover:r,nodeColorPressed:i,nodeColorActive:a,arrowColor:o,loadingColor:s,nodeTextColor:c,nodeTextColorDisabled:l,dropMarkColor:u,nodeWrapperPadding:d,nodeHeight:f,lineHeight:p,lineColor:h}}=m.value,g=se(d,`top`),_=se(d,`bottom`),v=he(pe(f)-pe(g)-pe(_));return{"--n-arrow-color":o,"--n-loading-color":s,"--n-bezier":e,"--n-font-size":t,"--n-node-border-radius":n,"--n-node-color-active":a,"--n-node-color-hover":r,"--n-node-color-pressed":i,"--n-node-text-color":c,"--n-node-text-color-disabled":l,"--n-drop-mark-color":u,"--n-node-wrapper-padding":d,"--n-line-offset-top":`-${g}`,"--n-line-offset-bottom":`-${_}`,"--n-node-content-height":v,"--n-line-height":p,"--n-line-color":h}}),at=i?ae(`tree`,void 0,it,t):void 0;return Object.assign(Object.assign({},nt),{mergedClsPrefix:n,mergedTheme:m,mergedRenderEmpty:g,rtlEnabled:p,fNodes:Ce,aip:X,selfElRef:v,virtualListInstRef:b,scrollbarInstRef:y,handleFocusout:et,handleDragLeaveTree:qe,handleScroll:Qe,getScrollContainer:x,getScrollContent:S,handleAfterEnter:Te,handleResize:$e,cssVars:i?void 0:it,themeClass:at?.themeClass,onRender:at?.onRender})},render(){var e;let{fNodes:t,internalRenderEmpty:n}=this;if(!t.length&&n)return n();let{mergedClsPrefix:r,blockNode:i,blockLine:a,draggable:o,disabled:s,ellipsis:c,internalFocusable:l,checkable:d,handleKeydown:f,rtlEnabled:p,handleFocusout:m,scrollbarProps:h}=this,g=l&&!s,_=g?`0`:void 0,v=[`${r}-tree`,p&&`${r}-tree--rtl`,d&&`${r}-tree--checkable`,(a||i)&&`${r}-tree--block-node`,a&&`${r}-tree--block-line`,c&&`${r}-tree--ellipsis`],y=e=>`__motion`in e?u(_t,{height:e.height,nodes:e.nodes,clsPrefix:r,mode:e.mode,onAfterEnter:this.handleAfterEnter}):u(gt,{key:e.key,tmNode:e,clsPrefix:r});if(this.virtualScroll){let{mergedTheme:e,internalScrollablePadding:n}=this,i=se(n||`0`);return u(fe,Object.assign({},h,{ref:`scrollbarInstRef`,onDragleave:o?this.handleDragLeaveTree:void 0,container:this.getScrollContainer,content:this.getScrollContent,class:v,theme:e.peers.Scrollbar,themeOverrides:e.peerOverrides.Scrollbar,tabindex:_,onKeydown:g?f:void 0,onFocusout:g?m:void 0}),{default:()=>{var n;return(n=this.onRender)==null||n.call(this),t.length?u(Te,{ref:`virtualListInstRef`,items:this.fNodes,itemSize:pe(e.self.nodeHeight),ignoreItemResize:this.aip,paddingTop:i.top,paddingBottom:i.bottom,class:this.themeClass,style:[this.cssVars,{paddingLeft:i.left,paddingRight:i.right}],onScroll:this.handleScroll,onResize:this.handleResize,showScrollbar:!1,itemResizable:!0},{default:({item:e})=>y(e)}):A(this.$slots.empty,()=>[this.mergedRenderEmpty?.call(this)||u(ye,{class:`${r}-tree__empty`,theme:this.mergedTheme.peers.Empty,themeOverrides:this.mergedTheme.peerOverrides.Empty})])}})}let{internalScrollable:b}=this;return v.push(this.themeClass),(e=this.onRender)==null||e.call(this),b?u(fe,Object.assign({},h,{class:v,tabindex:_,onKeydown:g?f:void 0,onFocusout:g?m:void 0,style:this.cssVars,contentStyle:{padding:this.internalScrollablePadding}}),{default:()=>u(`div`,{onDragleave:o?this.handleDragLeaveTree:void 0,ref:`selfElRef`},this.fNodes.map(y))}):u(`div`,{class:v,tabindex:_,ref:`selfElRef`,style:this.cssVars,onKeydown:g?f:void 0,onFocusout:g?m:void 0,onDragleave:o?this.handleDragLeaveTree:void 0},t.length?t.map(y):A(this.$slots.empty,()=>[this.mergedRenderEmpty?.call(this)||u(ye,{class:`${r}-tree__empty`,theme:this.mergedTheme.peers.Empty,themeOverrides:this.mergedTheme.peerOverrides.Empty})]))}});export{Fe as a,Ge as i,Qe as n,ke as o,Je as r,Ct as t};