# PartitionLabels

You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.

Return a list of integers representing the size of these parts.
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.

ababcbacadefegdehijhklij
012345678901234567890123
0        10        20

?=first-last
a=0-8
b=1-5
c=4-7
d=9-14
e=10-15
f=11-11
g=13-13
h=16-19
i=17-22
j=18-23
k=20-20
l=21-21

first=9
last=14

ababcbaca defegde hijhklij 