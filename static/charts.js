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