"""Calculate flight path"""
# You are going on a one-way indirect flight trip at includes an unknown very
# large number of transfers
#
# You are not stopping twice in the same airport.
# You have 1 ticket for each part of your trip.
# Each ticket contains a src and dst airport.
# All the tickets you have are randomly sorted.
# You forgot the original departure airport and your final destination.
#
# Write a function to reconstruct your trip with minimum big-O complexity.
#
# original:
# JFK -> ORD
# BOS -> MIA
# SFO -> BOS
# MIA -> LAX
# ORD -> SFO

def determine_origin(legs):
    """Find the origin of trip, it'll be the one that's never a destination"""
    for src in legs:
        if src not in legs.values():
            return src

def display_trip(origin, legs):
    """
    start with source = origin
    get source's destination
    print source and destination as the leg
    set source = destination and start over
    """
    src = origin
    while src in legs:
        print "%s -> %s" % (src, legs[src])
        src = legs[src]

def main():
    """Main, set up legs variable get origin and display trip"""
    legs = {
        'BOS': 'MIA',
        'SFO': 'BOS',
        'MIA': 'LAX',
        'ORD': 'SFO',
        'LAX': 'MCO',
        'JFK': 'ORD',
        'MCO': 'CHI',
        'CHI': 'DET',
        'DET': 'NYC',
        'NYC': 'LGD',
        'LGD': 'MSR',
    }
    count = 0
    origin = determine_origin(legs)
    display_trip(origin, legs)

if __name__ == "__main__":
    main()
