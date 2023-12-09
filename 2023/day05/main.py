import math


class Range:
    def __init__(self, dst_start, src_start, distance):
        self.dst_start = dst_start
        self.src_start = src_start
        self.distance = distance

    def is_src_in_range(self, src):
        return self.src_start <= src < self.src_start + self.distance

    def get_dst(self, val):
        return self.dst_start + val - self.src_start


def populate_ranges(ranges: list[Range], lines):
    while len(lines) > 0 and len(lines[0]) > 0:
        if "map" in lines[0]:
            lines = lines[1:]
        else:
            dst, src, dist = list(map(int, lines[0].split(" ")))
            ranges.append(Range(dst, src, dist))
            lines = lines[1:]
    return lines[1:]


def get_mapping(ranges: list[Range], val):
    for r in ranges:
        if r.is_src_in_range(val):
            return r.get_dst(val)
    return val


with open("input") as f:
    dat = f.read().strip("\n")

lines = dat.split("\n")
seeds = list(map(int, lines[0].split(":")[1].strip(" ").split(" ")))
# seeds_pairs = [(seeds[i], seeds[i + 1]) for i in range(0, len(seeds), 2)]
lines = lines[2:]

seed_to_soil = []
soil_to_fertilizer = []
fertilizer_to_water = []
water_to_light = []
light_to_temperature = []
temperature_to_humidity = []
humidity_to_location = []

lines = populate_ranges(seed_to_soil, lines)
lines = populate_ranges(soil_to_fertilizer, lines)
lines = populate_ranges(fertilizer_to_water, lines)
lines = populate_ranges(water_to_light, lines)
lines = populate_ranges(light_to_temperature, lines)
lines = populate_ranges(temperature_to_humidity, lines)
lines = populate_ranges(humidity_to_location, lines)

min_location = math.inf
for seed in seeds:
    soil = get_mapping(seed_to_soil, seed)
    fertilizer = get_mapping(soil_to_fertilizer, soil)
    water = get_mapping(fertilizer_to_water, fertilizer)
    light = get_mapping(water_to_light, water)
    temperature = get_mapping(light_to_temperature, light)
    humidity = get_mapping(temperature_to_humidity, temperature)
    location = get_mapping(humidity_to_location, humidity)
    min_location = min(min_location, location)

print(min_location)
