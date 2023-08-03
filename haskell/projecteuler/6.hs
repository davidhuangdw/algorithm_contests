
main = do
  let p6 = sum [1..100]^2 - sum (map (^2) [1..100])
  print p6
