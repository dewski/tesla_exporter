# tesla_exporter

<a href="https://heroku.com/deploy?template=https://github.com/dewski/tesla_exporter&env[TESLA_CLIENT_ID]=FINDME&env[TESLA_CLIENT_SECRET]=FINDME&env[TESLA_EMAIL]=account@email.com&env[TESLA_PASSWORD]=your_password_use_a_password_manager">
  <img src="https://www.herokucdn.com/deploy/button.svg" alt="Deploy">
</a>

This is a simple Prometheus exporter for your Tesla vehicles. It will display all vehicles in your account tagged accordingly using labels:

```
# HELP tesla_battery_level Vehicle charge readings
# TYPE tesla_battery_level gauge
tesla_battery_level{name="dewski",vin="1C6RR6KTXES448871"} 83
tesla_battery_level{name="otherdewski",vin="WBAWV53529PF81490"} 83
# HELP tesla_battery_level_usable Vehicle charge readings
# TYPE tesla_battery_level_usable gauge
tesla_battery_level_usable{name="dewski",vin="1C6RR6KTXES448871"} 82
# HELP tesla_battery_range Vehicle range readings
# TYPE tesla_battery_range gauge
tesla_battery_range{name="dewski",vin="1C6RR6KTXES448871"} 255.71
# HELP tesla_battery_range_estimaated Vehicle range readings
# TYPE tesla_battery_range_estimaated gauge
tesla_battery_range_estimaated{name="dewski",vin="1C6RR6KTXES448871"} 168.21
# HELP tesla_battery_range_ideal Vehicle range readings
# TYPE tesla_battery_range_ideal gauge
tesla_battery_range_ideal{name="dewski",vin="1C6RR6KTXES448871"} 319.63
# HELP tesla_odometer Vehicle odometer reading
# TYPE tesla_odometer gauge
tesla_odometer{name="dewski",vin="1C6RR6KTXES448871"} 5019.4068
# HELP tesla_temperature_driver Vehicle temperature readings
# TYPE tesla_temperature_driver gauge
tesla_temperature_driver{format="c",name="dewski",vin="1C6RR6KTXES448871"} 25
tesla_temperature_driver{format="f",name="dewski",vin="1C6RR6KTXES448871"} 77
# HELP tesla_temperature_inside Vehicle temperature readings
# TYPE tesla_temperature_inside gauge
tesla_temperature_inside{format="c",name="dewski",vin="1C6RR6KTXES448871"} 1
tesla_temperature_inside{format="f",name="dewski",vin="1C6RR6KTXES448871"} 33.79999923706055
# HELP tesla_temperature_outside Vehicle temperature readings
# TYPE tesla_temperature_outside gauge
tesla_temperature_outside{format="c",name="dewski",vin="1C6RR6KTXES448871"} -5.5
tesla_temperature_outside{format="f",name="dewski",vin="1C6RR6KTXES448871"} 22.100000381469727
# HELP tesla_temperature_passenger Vehicle temperature readings
# TYPE tesla_temperature_passenger gauge
tesla_temperature_passenger{format="c",name="dewski",vin="1C6RR6KTXES448871"} 25
tesla_temperature_passenger{format="f",name="dewski",vin="1C6RR6KTXES448871"} 77
```

## Copyright

Copyright © 2019 Garrett Bjerkhoel. See [MIT-LICENSE](/dewski/tesla_exporter/blob/master/MIT-LICENSE) for details.
