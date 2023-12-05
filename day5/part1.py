# todays problem was fun to figure out. think I identified a small shortcut for part 1 and excited to see how part 2 is going to make this more difficult

# read in the seeds and maps
import os


curr_dir = os.path.dirname(os.path.abspath(__file__))
os.chdir(curr_dir)
with open("input.txt") as f:
    lines = f.readlines()

index = 0

seeds = []
seed_line = lines[index].split(': ')[1].replace('\n', '').split(' ')
seeds = [int(s) for s in seed_line if s != '']
print("Seeds:", seeds)
index += 3

# read in soil map
all_soils = False
seed_to_soil = {}
while not all_soils:
    if lines[index] == '\n':
        all_soils = True
        index += 2
        break

    line = lines[index].replace('\n', '').split(' ')
    dest = int(line[0])
    lookup = int(line[1])
    r = int(line[2])
    seed_to_soil[lookup] = r, dest-lookup
    index += 1

# apply soil map to seeds
soils = []
for seed in seeds:
    soil = seed
    for k, v in seed_to_soil.items():
        if seed >= k and seed <= k+v[0]:
            soil += v[1]
            break
    soils.append(soil)
# free up memory
del seed_to_soil
print("Soils:", soils)

# read in fertilizer map
all_fertilizers = False
soil_to_fertilizer = {}
while not all_fertilizers:
    if lines[index] == '\n':
        all_fertilizers = True
        index += 2
        break
    line = lines[index].replace('\n', '').split(' ')
    dest = int(line[0])
    lookup = int(line[1])
    r = int(line[2])
    soil_to_fertilizer[lookup] = r, dest-lookup
    index += 1

# apply fertilizer map to soils
fertilizers = []
for soil in soils:
    fertilizer = soil
    for k, v in soil_to_fertilizer.items():
        if soil >= k and soil <= k+v[0]:
            fertilizer += v[1]
            break
    fertilizers.append(fertilizer)
# free up memory
del soil_to_fertilizer
print("Fertilizers:", fertilizers)

# read in water map
all_waters = False
fertilizer_to_water = {}
while not all_waters:
    if lines[index] == '\n':
        all_waters = True
        index += 2
        break
    line = lines[index].replace('\n', '').split(' ')
    dest = int(line[0])
    lookup = int(line[1])
    r = int(line[2])
    fertilizer_to_water[lookup] = r, dest-lookup
    index += 1

# apply water map to fertilizers
waters = []
for fertilizer in fertilizers:
    water = fertilizer
    for k, v in fertilizer_to_water.items():
        if fertilizer >= k and fertilizer <= k+v[0]:
            water += v[1]
            break
    waters.append(water)
# free up memory
del fertilizer_to_water
print("Waters:", waters)

# read in light map
all_lights = False
water_to_light = {}
while not all_lights:
    if lines[index] == '\n':
        all_lights = True
        index += 2
        break
    line = lines[index].replace('\n', '').split(' ')
    dest = int(line[0])
    lookup = int(line[1])
    r = int(line[2])
    water_to_light[lookup] = r, dest-lookup
    index += 1
# apply light map to waters
lights = []
for water in waters:
    light = water
    for k, v in water_to_light.items():
        if water >= k and water <= k+v[0]:
            light += v[1]
            break
    lights.append(light)
# free up memory
del water_to_light
print("Lights:", lights)

# read in temperature map
all_temperatures = False
light_to_temperature = {}
while not all_temperatures:
    if lines[index] == '\n':
        all_temperatures = True
        index += 2
        break
    line = lines[index].replace('\n', '').split(' ')
    dest = int(line[0])
    lookup = int(line[1])
    r = int(line[2])
    light_to_temperature[lookup] = r, dest-lookup
    index += 1
# apply temperature map to lights
temperatures = []
for light in lights:
    temperature = light
    for k, v in light_to_temperature.items():
        if light >= k and light <= k+v[0]:
            temperature += v[1]
            break
    temperatures.append(temperature)
# free up memory
del light_to_temperature
print("Temperatures:", temperatures)

# read in humidity map
all_humidities = False
temperature_to_humidity = {}
while not all_humidities:
    if lines[index] == '\n':
        all_humidities = True
        index += 2
        break
    line = lines[index].replace('\n', '').split(' ')
    dest = int(line[0])
    lookup = int(line[1])
    r = int(line[2])
    temperature_to_humidity[lookup] = r, dest-lookup
    index += 1
# apply humidity map to temperatures
humidities = []
for temperature in temperatures:
    humidity = temperature
    for k, v in temperature_to_humidity.items():
        if temperature >= k and temperature <= k+v[0]:
            humidity += v[1]
            break
    humidities.append(humidity)
# free up memory
del temperature_to_humidity
print("Humidities:", humidities)

# read in location map
humidity_to_location = {}
while index < len(lines):
    line = lines[index].replace('\n', '').split(' ')
    dest = int(line[0])
    lookup = int(line[1])
    r = int(line[2])
    humidity_to_location[lookup] = r, dest-lookup
    index += 1

# apply location map to humidities
locations = []
for humidity in humidities:
    location = humidity
    for k, v in humidity_to_location.items():
        if humidity >= k and humidity <= k+v[0]:
            location += v[1]
            break
    locations.append(location)
# free up memory
del humidity_to_location
print("Locations:", locations)

# print the lowest location
print("Lowest Location:", min(locations))

# debug
for i in range(len(seeds)):
    print(seeds[i], soils[i], fertilizers[i], waters[i], lights[i], temperatures[i], humidities[i], locations[i])


# first number is destination
# second number is lookup
# third number is the range

# walk each seed through the maps
    # check if the seed is in one of the (lookup to lookup+range) seed-to-soil maps
        # if it is, add the difference of (destination-lookup) to the location
        # if it is not, the soil is the same as the seed
    # check if the soil is in one of the (lookup to lookup+range) soil-to-fertilizer maps
        # if it is, add the difference of (destination-lookup) to the location
        # if it is not, the fertilizer is the same as the soil
    # check if the fertilizer is in one of the (lookup to lookup+range) fertilizer-to-water maps
        # if it is, add the difference of (destination-lookup) to the location
        # if it is not, the water is the same as the fertilizer
    # check if the water is in one of the (lookup to lookup+range) water-to-light maps
        # if it is, add the difference of (destination-lookup) to the location
        # if it is not, the light is the same as the water
    # check if the light is in one of the (lookup to lookup+range) light-to-temperature maps
        # if it is, add the difference of (destination-lookup) to the location
        # if it is not, the temperature is the same as the light
    # check if the temperature is in one of the (lookup to lookup+range) temperature-to-humidity maps
        # if it is, add the difference of (destination-lookup) to the location
        # if it is not, the humidity is the same as the temperature
    # check if the humidity is in one of the (lookup to lookup+range) humidity-to-location maps
        # if it is, add the difference of (destination-lookup) to the location
        # if it is not, the location is the same as the humidity

    # record the location of the seed in a dictionary with the seed as the key and the location as the value    

# after all seeds have been walked through the maps, print the lowest location