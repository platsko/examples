var APP = APP || {};

(function (app, $)
{
    app.getWeekDay = function (day) {
        var wd = '';
        switch (day) {
            case 0:
                wd = 'Sunday';
                break;
            case 1:
                wd = 'Monday';
                break;
            case 2:
                wd = 'Tuesday';
                break;
            case 3:
                wd = 'Wednesday';
                break;
            case 4:
                wd = 'Thursday';
                break;
            case 5:
                wd = 'Friday';
                break;
            case 6:
                wd = 'Saturday';
                break;
        }
        return wd;
    };

    app.formatDate = function (date, format) {
        if (!date) {
            return date;
        }
        var val = {
            d: date.getUTCDate(),
            m: date.getUTCMonth() + 1,
            Y: date.getUTCFullYear(),
            y: date.getUTCFullYear().toString().substring(2),
            H: date.getUTCHours(),
            h: date.getUTCHours(),
            i: date.getUTCMinutes(),
            s: date.getUTCSeconds()
        };

        format = {
            parts: format.split(/\W+/),
            separator: format.match(/[.\/\-\s:].*?/g)
        };

        if (!format.separator || !format.parts || format.parts.length < 1) {
            throw new Error('Invalid date format');
        }

        val.h = (val.h > 12) ? val.h - 12 : val.h;
        // prepend zero
        val.d = (val.d < 10 ? '0' : '') + val.d;
        val.m = (val.m < 10 ? '0' : '') + val.m;
        val.H = (val.H < 10 ? '0' : '') + val.H;
        val.h = (val.h < 10 ? '0' : '') + val.h;
        val.i = (val.i < 10 ? '0' : '') + val.i;
        val.s = (val.d < 10 ? '0' : '') + val.s;

        date = [];
        for (var i = 0, cnt = format.parts.length; i < cnt; i++) {
            date.push(val[format.parts[i]] + (format.separator[i] || ''));
        }

        return date.join('');
    };

    /** Charts Drawing */
    app.chart = function (config) {
        var chart, id, data, dates, legend, summary, columns, fraction, public_methods;
        chart = config.chart || {};
        id = 'chart' + chart.key;
        data = config.data;
        dates = d3.keys(data);
        legend = config.chart.legend;
        summary = {
            key: chart.key,
            title: chart.title
        };
        columns = [];
        fraction = config.chart.fraction || 0;
        public_methods = {
            pieChart: function () {
                return false;
            },
            drawPieChart: function () {
                return false;
            },
            drawPieChartTotal: function () {
                return false;
            },
            drawStackedBarChart: function () {
                return false;
            }
        };

        if (data.length === 0) {
            return public_methods;
        }

        public_methods.getItems = function (_data, filter) {
            var _items = [];
            Object.keys(_data).forEach(function (key) {
                var y0 = 0,
                    item = _data[key],
                    skip = ['count', 'key', 'legend', 'total'];
                item.count = [];
                Object.keys(item)
                    .filter(function (name) {
                        return skip.indexOf(name) === -1 && (!filter || name === filter);
                    })
                    .map(function (name) {
                        var _num = 0;
                        item.legend = name;
                        if (item[name] instanceof Object) {
                            for (var _key in item[name]) {
                                if (columns.indexOf(_key) === -1) {
                                    columns.push(_key);
                                }
                                _num = parseFloat(item[name][_key]);
                                item.count.push({
                                    name: name,
                                    column: _key,
                                    y0: y0,
                                    y1: (y0 += +_num).toFixed(fraction) * 1
                                });
                            }
                        } else {
                            if (columns.indexOf(key) === -1) {
                                columns.push(key);
                            }
                            _num = parseFloat(item[name]);
                            item.count.push({
                                name: name,
                                column: key,
                                y0: y0,
                                y1: (y0 += +_num).toFixed(fraction) * 1
                            });
                        }
                    });
                if (item.count.length) {
                    var _count = item.count[item.count.length - 1].y1;
                    item.key = key;
                    item.total = parseFloat(_count).toFixed(fraction) * 1;
                    _items.push(item);
                }
            });
            return _items;
        };

        var margin = {
                top: 350,
                right: 50,
                bottom: 100,
                left: 50
            },

            width = ($('#' + id).parent().innerWidth() || $('#_' + id).innerWidth()) - margin.left - margin.right,
            height = 700 - margin.top - margin.bottom,

            items = public_methods.getItems(data),

            y = d3.scale.linear()
                .domain([0, d3.max(items, function (d) {
                    return d.total;
                })])
                .rangeRound([height, 0]),
            yAxis = d3.svg.axis().scale(y).orient('left').tickFormat(d3.format('s')),

            x = d3.scale.ordinal().domain(dates).rangeBands([0, width]),
            xAxis = d3.svg.axis().scale(x.rangeBands([0, width])).tickValues(x.domain().filter(function (val, i, scope) {
                return i === 0 || i % Math.ceil(scope.length / 100) === 0;
            })).orient('bottom'),

            color = d3.scale.ordinal()
                .range([
                    'rgb(15,150,25)',
                    'rgb(50,100,200)',
                    'rgb(240,0,16)',
                    'rgb(150,0,150)',
                    'rgb(0,0,165)',
                    'rgb(0,128,128)',
                    'rgb(0,165,255)',
                    'rgb(255,155,0)',
                    'rgb(220,60,20)',
                    'rgb(128,128,128)',
                    'rgb(192,192,192)'
                ])
                .domain(legend || []),

            cdata = color.domain(),

            container = d3.select('#_' + id) || d3.select('#' + id),

            svg = container.append('svg')
                .attr('width', width + margin.left + margin.right)
                .attr('height', (height + margin.bottom + (cdata.length > 1 ? margin.top : 10)))
                .append('g')
                .attr('transform', 'translate(' + margin.left + ',' + (cdata.length > 1 ? margin.top : 10) + ')');

        if (config.sort) {
            items.sort(function (a, b) {
                return b.total - a.total;
            });
        }

        svg.append('g').attr('id', chart.key);

        svg.append('g')
            .attr('class', 'x axis')
            .attr('transform', 'translate(1,' + height + ')')
            .call(xAxis)
            .append('text')
            .attr('x', width - 25)
            .attr('dy', 100)
            .style('text-anchor', 'end')
            .text('Period');

        svg.selectAll('.x.axis .tick line')
            .attr('x1', -1)
            .attr('x2', -1);

        svg.selectAll('.x.axis .tick text')
            .attr('y', -5)
            .attr('x', -35)
            .attr('transform', 'rotate(-90)');

        svg.append('g')
            .attr('class', 'y axis')
            .call(yAxis)
            .append('text')
            .attr('transform', 'rotate(-90)')
            .attr('dy', -40)
            .style('text-anchor', 'end')
            .text(chart.title || chart.key.replace('_', ' '));

        var _idTimeout,
            pie_chart = {},
            state = svg.selectAll('.state')
                .data(items)
                .enter()
                .append('g')
                .attr('class', 'g')
                .attr('transform', function (d) {
                    return 'translate(' + x(d.key) + ',0)';
                });

        public_methods.drawStackedBarChart = function (name) {
            var _tip = d3.tip(),
                _items = public_methods.getItems(data, name);
            if (legend instanceof Array) {
                legend.forEach(function (val) {
                    summary[val] = 0;
                });
            }

            summary.total = _items.reduce(function (sum, item) {
                sum += item.count.reduce(function (_sum, _count) {
                    var _cnt = _count.y1 - _count.y0;
                    summary[_count.name] += _cnt;
                    return _sum + _cnt;
                }, 0);
                return sum;
            }, 0);

            svg.call(
                _tip
                    .attr('class', 'd3-tip')
                    .offset([-10, 0])
                    .html(function (d) {
                        return '<strong>' + d.key + '</strong>' +
                            '<span>&nbsp;&nbsp;[&nbsp;' + parseFloat(d.total).toFixed(fraction) + '&nbsp;]&nbsp;&nbsp;</span>' +
                            '<span style="font-style: italic;">' + (Math.round(d.total / summary.total * 1000) / 10) + '%</span>';
                    })
            );

            y.domain([0, d3.max(_items, function (d) {
                return d.total;
            })]).rangeRound([height, 0]);

            svg.select('.y.axis')
                .transition().duration(200)
                .call(yAxis);

            state.selectAll('rect')
                .transition().duration(100).ease('circleout')
                .remove();

            state.selectAll('rect')
                .data(_items)
                .exit()
                .data(function (d) {
                    return d.count;
                })
                .enter()
                .append('rect')
                .transition().duration(100)
                .attr('width', Math.floor(x.rangeExtent()[1] / x.range().length))
                .attr('y', function (d) {
                    return y(d.y1);
                })
                .attr('height', function (d) {
                    return y(d.y0) - y(d.y1);
                })
                .style('fill', function (d) {
                    return color(d.name);
                });

            state.on('mouseenter', name || cdata.length === 1 ? _tip.show : public_methods.drawPieChart);
            state.on('mouseleave', name || cdata.length === 1 ? _tip.hide : public_methods.drawPieChartTotal);
        };

        public_methods.drawPieChart = function (data) {
            if (pie_chart.key && pie_chart.key === data.key) {
                return true;
            }
            if (_idTimeout) {
                clearTimeout(_idTimeout);
            }
            _idTimeout = setTimeout(function () {
                if (pie_chart.key) {
                    pie_chart.destroy();
                }
                pie_chart = public_methods.pieChart(cdata.map(function (type) {
                    var _num = data[type] instanceof Object
                        ? d3.values(data[type]).reduce(function (_prev, _curr) {
                        return parseFloat(_prev) + parseFloat(_curr);
                    }) || 0
                        : parseFloat(data[type]) || 0;
                    return {
                        type: type,
                        label: type + ' [ ' + (_num.toFixed(fraction)) + ' ]',
                        value: (_num * 1) || 0,
                        color: color(type)
                    };
                }), {
                    key: '_pie' + data.key.replace(' ', '_'),
                    total: parseFloat(data.total).toFixed(fraction),
                    title: data.key || chart.title || chart.key.replace('_', ' '),
                });
            }, 800);
            return true;
        };

        public_methods.drawPieChartTotal = function () {
            summary.key = chart.title || chart.key.replace('_', ' ');
            return public_methods.drawPieChart(summary);
        };

        public_methods.pieChart = function (data, conf) {
            var _chart,
                _svg = svg.select('#' + chart.key),
                _width = 325,
                _conf = {
                    'header': {
                        'title': {
                            'text': conf.title,
                            'fontSize': 16,
                            'font': 'courier',
                            'color': '#999999',
                        },
                        'subtitle': {
                            'text': (chart.labelTotal || 'Total') + ': ' + conf.total,
                            'fontSize': 12,
                            'font': 'courier'
                        },
                        'location': 'pie-center',
                        'titleSubtitlePadding': 10
                    },
                    'size': {
                        'canvasWidth': _width,
                        'pieOuterRadius': '95%',
                        'pieInnerRadius': '65%'
                    },
                    'data': {
                        'sortOrder': 'none',
                        'content': data
                    },
                    'labels': {
                        'outer': {
                            'format': 'label-percentage1',
                            'pieDistance': 25
                        },
                        'inner': {
                            'format': 'none',
                            'hideWhenLessThanPercentage': 1
                        },
                        'mainLabel': {
                            'color': '#666666',
                            'fontSize': 11
                        },
                        'percentage': {
                            'color': '#666',
                            'decimalPlaces': 1
                        },
                        'value': {
                            'color': '#999',
                            'fontSize': 11
                        },
                        'lines': {
                            'enabled': true
                        },
                    },
                    'effects': {
                        'load': {
                            'speed': 500
                        },
                        'pullOutSegmentOnClick': {
                            'size': 8,
                            'speed': 300,
                            'effect': 'linear'
                        },
                        'highlightSegmentOnMouseover': false
                    },
                    'misc': {
                        'gradient': {
                            'enabled': false,
                            'percentage': 95,
                            'color': '#fff'
                        }
                    },
                    'callbacks': {
                        'onload': null,
                        'onMouseoverSegment': function (seg) {
                            seg.segment.style.cursor = 'pointer';
                            seg.segment.style.opacity = 0.8;
                        },
                        'onMouseoutSegment': function (seg) {
                            seg.segment.style.opacity = 1;
                        },
                        'onClickSegment': function (seg) {
                            var name = seg.expanded ? null : seg.data.type,
                                leg = keywords[0].filter(function (_d) {
                                    return _d.textContent === name;
                                });
                            public_methods.drawStackedBarChart(name);
                            keywords.attr('class', 'legend').attr('filtered', null);
                            if (name) {
                                keywords.attr('class', 'legend mute');
                                d3.select(leg[0]).attr('class', 'legend filtered').attr('filtered', 1);
                            }
                        }
                    }
                };
            _svg.attr('transform', 'translate(' + (Math.round(width / 2 - _width / 2)) + ',-' + Math.round(_width + 100) + ')');
            _chart = new d3pie(chart.key, _conf);
            _chart.key = conf.key || new Date().getTime();
            _svg.selectAll('svg').attr('style', 'overflow:visible;');
            return _chart;
        };

        var keywords = svg.selectAll('.legend').data(cdata);
        if (keywords && keywords[0] && keywords[0].length > 1) {
            keywords.enter()
                .append('g')
                .attr('class', 'legend')
                .attr('transform', function (d, i) {
                    return 'translate(0,' + (i * 20 - 250) + ')';
                });
            keywords.append('rect')
                .attr('x', width - 18)
                .attr('width', 18)
                .attr('height', 18)
                .style('fill', color);
            keywords.append('text')
                .attr('x', width - 24)
                .attr('y', 9)
                .attr('dy', '.35em')
                .style('text-anchor', 'end')
                .text(function (d) {
                    return d;
                });
            keywords.on('click', function (val, idx) {
                var seg = -1,
                    leg = d3.select(keywords[0][idx]),
                    val = leg.attr('filtered') ? null : val;
                pie_chart.closeSegment();
                public_methods.drawStackedBarChart(val);
                keywords.attr('class', 'legend').attr('filtered', null);
                if (val) {
                    pie_chart.options.data.content.forEach(function (d, i) {
                        var found = d.type == val;
                        if (found) {
                            seg = i;
                        }
                        return !found;
                    });
                    keywords.attr('class', 'legend mute');
                    d3.select(keywords[0][idx]).attr('class', 'legend filtered').attr('filtered', 1);
                    pie_chart.openSegment(seg);
                }
            });
        }

        return public_methods;
    };

    // Base64 functions
    app.Base64 = new (function () {
        var self = this,
            public_methods = {
                // public_methods method encoding
                encode: function (input) {
                    var chr1, chr2, chr3, enc1, enc2, enc3, enc4,
                        i = 0, output = '';
                    input = self.utf8_encode(input);
                    while (i < input.length) {
                        chr1 = input.charCodeAt(i++);
                        chr2 = input.charCodeAt(i++);
                        chr3 = input.charCodeAt(i++);
                        enc1 = chr1 >> 2;
                        enc2 = ((chr1 & 3) << 4) | (chr2 >> 4);
                        enc3 = ((chr2 & 15) << 2) | (chr3 >> 6);
                        enc4 = chr3 & 63;
                        if (isNaN(chr2)) {
                            enc3 = enc4 = 64;
                        } else if (isNaN(chr3)) {
                            enc4 = 64;
                        }
                        output = output +
                            self.keyStr.charAt(enc1) + self.keyStr.charAt(enc2) +
                            self.keyStr.charAt(enc3) + self.keyStr.charAt(enc4);
                    }
                    return output;
                },
                // public_methods method decoding
                decode: function (input) {
                    var chr1, chr2, chr3, enc1, enc2, enc3, enc4,
                        i = 0, output = '';
                    input = input.replace(/[^A-Za-z0-9\+\/\=]/g, "");
                    while (i < input.length) {
                        enc1 = self.keyStr.indexOf(input.charAt(i++));
                        enc2 = self.keyStr.indexOf(input.charAt(i++));
                        enc3 = self.keyStr.indexOf(input.charAt(i++));
                        enc4 = self.keyStr.indexOf(input.charAt(i++));
                        chr1 = (enc1 << 2) | (enc2 >> 4);
                        chr2 = ((enc2 & 15) << 4) | (enc3 >> 2);
                        chr3 = ((enc3 & 3) << 6) | enc4;
                        output = output + String.fromCharCode(chr1);
                        if (enc3 !== 64) {
                            output = output + String.fromCharCode(chr2);
                        }
                        if (enc4 !== 64) {
                            output = output + String.fromCharCode(chr3);
                        }
                    }
                    output = self.utf8_decode(output);
                    return output;
                }
            };
        // private method for UTF-8 encoding
        this.keyStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
        this.utf8_encode = function (string) {
            string = string.replace(/\r\n/g, "\n");
            var utftext = '';
            for (var n = 0; n < string.length; n++) {
                var c = string.charCodeAt(n);
                if (c < 128) {
                    utftext += String.fromCharCode(c);
                }
                else if ((c > 127) && (c < 2048)) {
                    utftext += String.fromCharCode((c >> 6) | 192);
                    utftext += String.fromCharCode((c & 63) | 128);
                }
                else {
                    utftext += String.fromCharCode((c >> 12) | 224);
                    utftext += String.fromCharCode(((c >> 6) & 63) | 128);
                    utftext += String.fromCharCode((c & 63) | 128);
                }
            }
            return utftext;
        };
        // private method for UTF-8 decoding
        this.utf8_decode = function (utfString) {
            var i = 0, c = 0, c1 = 0, c2 = 0,
                string = '';
            while (i < utfString.length) {
                c = utfString.charCodeAt(i);
                if (c < 128) {
                    string += String.fromCharCode(c);
                    i++;
                }
                else if ((c > 191) && (c < 224)) {
                    c1 = utfString.charCodeAt(i + 1);
                    string += String.fromCharCode(((c & 31) << 6) | (c1 & 63));
                    i += 2;
                }
                else {
                    c1 = utfString.charCodeAt(i + 1);
                    c2 = utfString.charCodeAt(i + 2);
                    string += String.fromCharCode(((c & 15) << 12) | ((c1 & 63) << 6) | (c2 & 63));
                    i += 3;
                }
            }
            return string;
        };
        return public_methods;
    }
)();
}(APP, jQuery));
