from dataclasses import dataclass

@dataclass(frozen=True)
class Coordinate:
    x: int
    y: int

    @staticmethod
    def l1_distance(c1: 'Coordinate', c2: 'Coordinate') -> int:
        return abs(c1.x - c2.x) + abs(c1.y - c2.y)


N = int(input())

t_prev = 0
c_prev = Coordinate(0, 0)

for _ in range(N):
    splited = input().split(' ')
    t = int(splited[0])
    c = Coordinate(int(splited[1]), int(splited[2]))

    t_now = t - t_prev
    distance = Coordinate.l1_distance(c, c_prev)

    if t_now < distance:
        print('No')
        exit(0)

    if t_now % 2 != distance % 2:
        print('No')
        exit(0)

    c_prev = c
    t_prev = t

print('Yes')
