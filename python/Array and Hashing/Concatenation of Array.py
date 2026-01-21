from typing import List

class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        
        nums += nums #appending list to an existing list
        return nums