# todays problem was fun to figure out. think I identified a small shortcut for part 1 and excited to see how part 2 is going to make this more difficult

# read in the seeds and maps
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