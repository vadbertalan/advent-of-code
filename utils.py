class Coord():
    def __init__(self, x, y):
        self.x = x
        self.y = y
    def __str__(self):
        return f'({self.x}, {self.y})'
    def __repr__(self):
        return f'({self.x}, {self.y})'
    def get_tuple(self):
        return self.x, self.y
    def __eq__(self, other):
        return self.x == other.x and self.y == other.y
