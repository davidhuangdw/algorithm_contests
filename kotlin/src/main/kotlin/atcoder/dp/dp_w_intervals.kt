package atcoder.dp

// idea: see solution -- choose iteration feature 'last 1'
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (N, M) = input_ints()
    val segs = (1..M).map { input_longs() }.sortedBy { it[0] }
    val dp = MutableList(N+1){0L}
//    fun contri(j: Int, i: Int): Long { return segs.filter{ it[0] in (j + 1)..i && i <= it[1]}.sumOf { it[2] } }
//    for(i in 1..N) {
////        dp[i] = (0 until i).maxOf { j -> dp[j] + contri(j, i) }
//        val ci = segs.filter{i in it[0]..it[1]}
//        var sum = ci.sumOf { it[2] }
//        var k = 0
//        dp[i] = (0 until i).maxOf { j ->
//            while(k<ci.size && ci[k][0] <= j) {
//                sum -= ci[k][2]
//                k++
//            }
//            dp[j] + sum
//        }
//    }

    class SegTree(val n: Int) { // 0..n
        private val tree = LongArray(4 * (n+1))
        private val lazy = LongArray(4 * (n+1))

        private fun push(i: Int, l: Int, r: Int) {
            if(lazy[i] == 0L) return
            tree[i] += lazy[i]
            if(l < r){
                lazy[i*2] += lazy[i]
                lazy[i*2 + 1] += lazy[i]
            }
            lazy[i] = 0L
        }

        fun add(l: Int, r: Int, v: Long)=add(1, 0, n, l, r, v)

        private fun add(i: Int, tl: Int, tr: Int, l: Int, r: Int, v: Long) {
            push(i, tl, tr) // bug: always need push -- because the parent use tree[i] without lazy[i]
            if(tr < l || r < tl || l > r) return
            if(l <= tl && tr <= r){
                lazy[i] += v
                push(i, tl, tr) // ensure i is pushed
            }else {
                val tm = (tl + tr) / 2
                add(i*2, tl, tm, l,  r, v)
                add(i*2+1, tm+1, tr, l, r, v)
                tree[i] = maxOf(tree[i*2], tree[i*2+1])
            }
        }

        fun max(l: Int, r: Int): Long = max(1, 0, n, l, r)

        private fun max(i: Int, tl: Int, tr: Int, l: Int, r: Int): Long {
            push(i, tl, tr)
            if(tr < l || r < tl || l > r) return Long.MIN_VALUE
            if(l <= tl && tr <= r){
                return tree[i]
            }
            val tm = (tl + tr) / 2
            return maxOf(
                max(i*2, tl, tm, l, r),
                max(i*2+1, tm+1, tr, l, r)
            )
        }

    }

    val tr = SegTree(N)
    val segsLeft = segs.groupBy { it[0].toInt() }
    val segsRight = segs.groupBy { it[1].toInt() }

    for(i in 1..N){
        for(s in segsLeft.getOrDefault(i, listOf())){
            val (l, _, v) = s
//            println("left: $s")
            tr.add(0, l.toInt()-1, v)
        }

        dp[i] = tr.max(0, i-1)
        tr.add(i, i, dp[i])
        for(s in segsRight.getOrDefault(i, listOf())){
            val (l, _, v) = s
//            println("right: $s")
            tr.add(0, l.toInt()-1, -v)
        }
    }

//    println(dp)
//    println(dp1)
    println(dp.maxOrNull()!!)
}
