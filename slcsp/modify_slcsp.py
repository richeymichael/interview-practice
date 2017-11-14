"""Second Lowest Cost Silver Plan Homework"""


METAL_LEVEL = "Silver"
ZIPS_FILE = "zips.csv"
PLANS_FILE = "plans.csv"
SLCSP_FILE = 'slcsp.csv'


def build_plan_table(planfile, level):
    """Create a plan table for all plans at a specified metal level

    example plans dict:
        {
            'FL': {
                2: [366.01, 330.9, 320.76, 298.24, 324.61, 294.87, ...],
                1: [326.65, 330.9, 356.11, 375.97, 309.77, 340.29, ...],
                ...
            }
            ...
        }
    """
    plans = {}
    header = True
    with open(planfile, 'r') as inputf:
        for line in inputf:
            if not header:
                line = line[:-1]
                (_, state, metal_level, rate, rate_area) = line.split(',')
                rate = float(rate)
                rate_area = int(rate_area)
                if metal_level == level:
                    if state in plans:
                        if rate_area in plans[state]:
                            if rate not in plans[state][rate_area]:
                                plans[state][rate_area].append(rate)
                        else:
                            plans[state][rate_area] = [rate]
                    else:
                        plans[state] = {
                            rate_area:  [rate]
                        }
            else:
                header = False
    return plans


def build_zip_table(zipfile):
    """Create a zips table for all zipcode mappings

    example zips dict:
        {
            '32839': {
                48: ['FL']
            },
            '42223': {
                2: ['KY'],
                4: ['TN']
            },
            '59221': {
                4: ['MT', 'ND']
            },
            ...
        }
    """
    zips = {}
    header = True
    with open(zipfile, 'r') as inputf:
        for line in inputf:
            if not header:
                line = line[:-1]
                (zipcode, state, _, _, rate_area) = line.split(',')
                rate_area = int(rate_area)
                if zipcode in zips:
                    if rate_area in zips[zipcode]:
                        if state not in zips[zipcode][rate_area]:
                            zips[zipcode][rate_area].append(state)
                    else:
                        zips[zipcode][rate_area] = [state]
                else:
                    zips[zipcode] = {
                        rate_area: [state]
                    }
            else:
                header = False
    return zips


def find_second_lowest(zips, plans, zipcode):
    """Find the second lowest rate and return it

    return None if a rate area has only one rate
    return None if a rate area has no plans
    return None if a state has no plans
    return None if a rate area has multiple states
    return None if a zip has multiple rate areas
    return None if a zip has multiple rate areas
    """
    if len(zips[zipcode]) == 1:
        rate_area = zips[zipcode].keys()[0]
        if len(zips[zipcode][rate_area]) == 1:
            state = zips[zipcode][rate_area][0]
            if state in plans:
                if rate_area in plans[state]:
                    rates = plans[state][rate_area][:]
                    if len(rates) > 1:
                        rates.remove(min(rates))
                        return min(rates)
                    else:
                        return None  # only 1 rate for area
                else:
                    return None  # no rates for area
            else:
                return None  # no plans for state
        else:
            return None  # multiple states with same rate area for zip
    else:
        return None  # multiple rate areas for zip


def main():
    """Main method

    Build a dict with the zips
    Build a dict with the plans
    Read in all the lines in the SLCSP file
    Overwrite the SLCSP file line by line
        calling find_second_lowest for the zipcode on each line
    """
    zips = build_zip_table(ZIPS_FILE)
    plans = build_plan_table(PLANS_FILE, METAL_LEVEL)
    with open(SLCSP_FILE, 'r') as inputf:
        lines = inputf.readlines()
    with open(SLCSP_FILE, 'w') as outputf:
        header = True
        for line in lines:
            if not header:
                (zipcode, _) = line.split(',')
                slcsp = find_second_lowest(zips, plans, zipcode)
                if slcsp:
                    outputf.write("{}, {}\n".format(zipcode, slcsp))
                else:
                    outputf.write(line)
            else:
                outputf.write(line)
                header = False


if __name__ == "__main__":
    main()
