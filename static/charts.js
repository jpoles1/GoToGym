function generateGauge(divID, gaugeValue){
    c3.generate({
        bindto: divID,
        data: {
            columns: [
                ['% Attendance', gaugeValue]
            ],
            type: 'gauge',
        },
        color: {
            pattern: ['#FF0000', '#F97600', '#F6C600', '#60B044'], // the three color levels for the percentage values.
            threshold: {
                values: [30, 60, 90, 100]
            }
        },
        size: {
            height: 180
        }
    });
}
function generateAttendancePlot(divID, plotData){
    c3.generate({
        bindto: divID,
        data: {
            x: 'Time',
            //March 2, 2018 at 10:08PM
            xFormat: '%B %-d, %Y at %-I:%M%p', // 'xFormat' can be used as custom format of 'x'
            columns: plotData
        },
        axis: {
            x: {
                type: 'timeseries',
                tick: {
                    rotate: 15,
                    format: "%H:%M:%S %Y-%m-%d"
                }
            }
        },
        zoom: {
            enabled: true
        }
    });
}