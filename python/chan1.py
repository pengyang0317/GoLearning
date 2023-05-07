
# 定义随机双色球函数

import random

def generate_random():
    # 生成双色球随机列表
    red_balls = [x for x in range(1,34)]
    selected_balls = random.sample(red_balls,6)
    selected_balls.sort()
    selected_balls.append(random.randint(1,16))
    return selected_balls

def display(balls):
    # 显示双色球列表
    for index, ball in enumerate(balls):
        if index == len(balls) - 1:
            print('|', end=' ')
        print('%02d' % ball, end=' ')
    print()

def main():
    # 主函数
    n = int(input('机选几注: '))
    for _ in range(n):
        display(generate_random())

if __name__ == '__main__':
    main()
