import json
import random
import os
import subprocess
import sys
import time


ITEMS = ['money', 'computer', 'shoe', 'dress', 'car', 'house', 'apple',
    'pencil', 'foot', 'star', 'planet', 'color', 'straw', 'battery',
    'controller', 'box', 'ram', 'harddrive', 'phone', 'toe', 'arm',
    'spanish', 'child', 'school', 'math', 'shorts', 'pizza', 'peroxide',
    'duster', 'light', 'table', 'up', 'cord', 'iron', 'bed', 'closet']
ITEM_COUNT = (5, 16)
ITEM_SCORE = (1, 10)
POOL_COUNT = 100000


def make_item():
    data = {}

    for i in range(ITEM_COUNT[0], ITEM_COUNT[1]):
        tag = random.choice(ITEMS)
        data[tag] = random.randint(ITEM_SCORE[0], ITEM_SCORE[1])

    return {
        'id': str(random.random()),
        'data': data,
    }


def run(source, pool, count=None):
    if count:
        pool = pool[:count]

    source = json.dumps(source)
    pool_len = len(pool)
    pool = json.dumps(pool)
    pool_file = '/tmp/gosinesim-{}.json'.format(pool_len)

    with open(pool_file, 'w') as j:
        j.write(pool)

    def runner(worker=False):
        worker_pool = ''

        if worker:
            worker_pool = 'using the worker pool'
            cmd = "./gosinesim --source='{}' --pool_file='{}' --worker=true > /dev/null 2>&1".format(
                source, pool_file)
        else:
            cmd = "./gosinesim --source='{}' --pool_file='{}' > /dev/null 2>&1".format(
                source, pool_file)

        print "\n"
        print '*' * 40
        print "\nComparing {} items {} \n".format(pool_len, worker_pool)

        for i in range(3):
            start = time.time()
            os.system(cmd)
            dif = time.time() - start

            print "\trun {} took {}".format(i + 1, dif)

    runner()
    runner(worker=True)


if __name__ == '__main__':
    print "*" * 60
    print "\nBuilding the pool of {} items".format(POOL_COUNT)

    start = time.time()
    source = make_item()
    pool = [make_item() for i in range(POOL_COUNT)]
    counts = [100, 1000, 5000, 10000, 25000, 50000, 75000, None]

    print "\ndone building the pool: {}".format(time.time() - start)

    [run(source, pool, c) for c in counts]

    print "\n"
    print "=" * 60
    print "\nfinished running: {}".format(time.time() - start)
