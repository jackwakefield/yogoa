/**
 * Copyright (c) 2014-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

function toValueGo(value) {
  return value.toString().replace('px','').replace('%','');
}

function toMethodName(value) {
  if (value.indexOf('%') >= 0){
    return 'Percent';
  } else if(value.indexOf('Auto') >= 0) {
    return 'Auto';
  }
  return '';
}

function toExportName(name) {
  name = name.replace(/(\_\w)/g, function(m) { return m[1].toUpperCase(); });
  if (name.length > 0) {
    name = name[0].toUpperCase() + name.substring(1);
  }
  return name;
}

var GoEmitter = function() {
  Emitter.call(this, 'go', '  ');
};

GoEmitter.prototype = Object.create(Emitter.prototype, {
  constructor:{value:GoEmitter},

  emitPrologue:{value:function() {}},

  emitTestPrologue:{value:function(name, experiments) {
    this.push('func Test' + toExportName(name) + '(t *testing.T) {');
    this.pushIndent();

    this.push('config := NewConfig()')
    this.push('defer config.Free()')
    for (var i in experiments) {
      this.push('config.SetExperimentalFeatureEnabled(Experiment' + experiments[i] +', true)');
    }
    this.push('');
  }},

  emitTestTreePrologue:{value:function(nodeName) {
    this.push(nodeName + ' := NewNodeWithConfig(config)');
    this.push('defer ' +nodeName + '.Free()')
  }},

  emitTestEpilogue:{value:function(experiments) {
    this.popIndent();
    this.push('}');
    this.push('');
  }},

  emitEpilogue:{value:function(lines) {}},

  AssertEQ:{value:function(v0, v1) {
    this.push('assert.EqualValues(t, ' + v0 + ', ' + v1 + ')');
  }},

  YGAlignAuto:{value:'AlignAuto'},
  YGAlignCenter:{value:'AlignCenter'},
  YGAlignFlexEnd:{value:'AlignFlexEnd'},
  YGAlignFlexStart:{value:'AlignFlexStart'},
  YGAlignStretch:{value:'AlignStretch'},
  YGAlignSpaceBetween:{value:'AlignSpaceBetween'},
  YGAlignSpaceAround:{value:'AlignSpaceAround'},
  YGAlignBaseline:{value:'AlignBaseline'},

  YGDirectionInherit:{value:'DirectionInherit'},
  YGDirectionLTR:{value:'DirectionLTR'},
  YGDirectionRTL:{value:'DirectionRTL'},

  YGEdgeBottom:{value:'EdgeBottom'},
  YGEdgeEnd:{value:'EdgeEnd'},
  YGEdgeLeft:{value:'EdgeLeft'},
  YGEdgeRight:{value:'EdgeRight'},
  YGEdgeStart:{value:'EdgeStart'},
  YGEdgeTop:{value:'EdgeTop'},

  YGFlexDirectionColumn:{value:'FlexDirectionColumn'},
  YGFlexDirectionColumnReverse:{value:'FlexDirectionColumnReverse'},
  YGFlexDirectionRow:{value:'FlexDirectionRow'},
  YGFlexDirectionRowReverse:{value:'FlexDirectionRowReverse'},

  YGJustifyCenter:{value:'JustifyCenter'},
  YGJustifyFlexEnd:{value:'JustifyFlexEnd'},
  YGJustifyFlexStart:{value:'JustifyFlexStart'},
  YGJustifySpaceAround:{value:'JustifySpaceAround'},
  YGJustifySpaceBetween:{value:'JustifySpaceBetween'},
  YGJustifySpaceEvenly:{value:'JustifySpaceEvenly'},

  YGOverflowHidden:{value:'OverflowHidden'},
  YGOverflowVisible:{value:'OverflowVisible'},

  YGPositionTypeAbsolute:{value:'PositionAbsolute'},
  YGPositionTypeRelative:{value:'PositionRelative'},

  YGUndefined:{value:'Undefined'},

  YGDisplayFlex:{value:'DisplayFlex'},
  YGDisplayNone:{value:'DisplayNone'},
  YGAuto:{value:'Auto'},


  YGWrapNoWrap:{value:'WrapNone'},
  YGWrapWrap:{value:'WrapWrap'},
  YGWrapWrapReverse:{value: 'WrapReverse'},

  YGNodeCalculateLayout:{value:function(node, dir, experiments) {
    this.push(node + '.Style().SetDirection(' + dir + ')');
    this.push(node + '.CalculateLayout(Undefined, Undefined, ' + dir +')');
  }},

  YGNodeInsertChild:{value:function(parentName, nodeName, index) {
    this.push(parentName + '.InsertChild(' + nodeName + ', ' + index + ')');
  }},

  YGNodeLayoutGetLeft:{value:function(nodeName) {
    return nodeName + '.Layout().Left()';
  }},

  YGNodeLayoutGetTop:{value:function(nodeName) {
    return nodeName + '.Layout().Top()';
  }},

  YGNodeLayoutGetWidth:{value:function(nodeName) {
    return nodeName + '.Layout().Width()';
  }},

  YGNodeLayoutGetHeight:{value:function(nodeName) {
    return nodeName + '.Layout().Height()';
  }},

  YGNodeStyleSetAlignContent:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetAlignContent(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetAlignItems:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetAlignItems(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetAlignSelf:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetAlignSelf(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetBorder:{value:function(nodeName, edge, value) {
    this.push(nodeName + '.Style().SetBorder(' + edge + ', ' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetDirection:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetDirection(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetDisplay:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetDisplay(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetFlexBasis:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetFlexBasis' + toMethodName(value) + '(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetFlexDirection:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetFlexDirection(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetFlexGrow:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetFlexGrow(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetFlexShrink:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetFlexShrink(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetFlexWrap:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetFlexWrap(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetHeight:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetHeight' + toMethodName(value) + '(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetJustifyContent:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetJustifyContent(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetMargin:{value:function(nodeName, edge, value) {
    var valueStr = toValueGo(value);
    if (valueStr != 'Auto') {
      valueStr = ', ' + valueStr + '';
    } else {
      valueStr = '';
    }

    this.push(nodeName + '.Style().SetMargin' + toMethodName(value) + '(' + edge + valueStr + ')');
  }},

  YGNodeStyleSetMaxHeight:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetMaxHeight' + toMethodName(value) + '(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetMaxWidth:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetMaxWidth' + toMethodName(value) + '(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetMinHeight:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetMinHeight' + toMethodName(value) + '(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetMinWidth:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetMinWidth' + toMethodName(value) + '(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetOverflow:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetOverflow(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetPadding:{value:function(nodeName, edge, value) {
    this.push(nodeName + '.Style().SetPadding' + toMethodName(value) + '(' + edge + ', ' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetPosition:{value:function(nodeName, edge, value) {
    this.push(nodeName + '.Style().SetPosition' + toMethodName(value) + '(' + edge + ', ' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetPositionType:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetPositionType(' + toValueGo(value) + ')');
  }},

  YGNodeStyleSetWidth:{value:function(nodeName, value) {
    this.push(nodeName + '.Style().SetWidth' + toMethodName(value) + '(' + toValueGo(value) + ')');
  }},
});