import sys

fn = sys.argv[1] if len(sys.argv) > 1 else "test"

ans = 0
for x in range(-100, 1000):
    for y in range(1000):
        ok = False
        vx = x
        max_y = 0
        vy = y
        dx = 0
        dy = 0
        for t in range(500):
            dx += vx
            dy += vy
            max_y = max(max_y, dy)
            if vx > 0:
                vx -= 1
            elif vx < 0:
                vx += 1
            vy -= 1
            if 139 <= dx <= 187 and -148 <= dy <= -89:
                ok = True
        if ok:
            print(x, y)
            ans = max(max_y, ans)

print(ans)
