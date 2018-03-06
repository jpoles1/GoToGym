var gaugeColorScale = ['#FF0000', '#F97600', '#F6C600', '#60B044']
function generateGauge(divID, gaugeValue) {
    c3.generate({
        bindto: divID + " .plot-box",
        data: {
            columns: [
                ['% Attendance', gaugeValue]
            ],
            type: 'gauge',
        },
        color: {
            pattern: gaugeColorScale,
            threshold: {
                values: [30, 60, 90, 100]
            }
        },
        size: {
            height: 180
        }
    });
}

function generateMonthlyMissedGauge(divID, missedCount, totalCount) {
    c3.generate({
        bindto: divID + " .plot-box",
        data: {
            columns: [
                ['Missed this Month', missedCount]
            ],
            type: 'gauge',
        },
        gauge: {
            max: totalCount,
            label: {
                format: function (value, ratio) {
                    return value;
                },
            }
        },
        color: {
            pattern: gaugeColorScale.reverse(), // the three color levels for the percentage values.
            threshold: {
                values: [30, 60, 90, 100]
            }
        },
        size: {
            height: 180
        }
    });
}

function generateAttendancePlot(divID, plotData) {
    c3.generate({
        bindto: divID + " .plot-box",
        data: {
            x: 'Time',
            //March 2, 2018 at 10:08PM
            xFormat: '%B %-d, %Y at %-I:%M%p', // 'xFormat' can be used as custom format of 'x'
            columns: plotData
        },
        padding: {
            left: 40,
            right: 40,
            top: 20
        },
        axis: {
            x: {
                type: 'timeseries',
                tick: {
                    rotate: 15,
                    format: "%H:%M:%S %m-%d-%Y"
                }
            }
        },
        zoom: {
            enabled: true
        }
    });
}
function generateDayOfWeekPlot(divID, dayOfWeekCounts){
    c3.generate({
        bindto: divID + " .plot-box",
        data: {
            json: dayOfWeekCounts,
            type : 'donut',
        },
        donut: {
            title: "Workout DoW"
        }
    });
}