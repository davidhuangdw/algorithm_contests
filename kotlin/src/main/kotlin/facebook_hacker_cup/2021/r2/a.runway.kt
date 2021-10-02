
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, M) = input_ints()
        fun to_count(arr: List<Int>) = arr.fold(hashMapOf<Int, Int>()){cnt, v->
            cnt[v] = cnt.getOrDefault(v, 0) + 1
            cnt
        }
        val rounds = List(N+1){to_count(input_ints())}
        val keep = rounds[0]
        var diff_sum = 0
        for((a, b) in rounds.zipWithNext()){
            for((v, c) in keep) {
                if(c > 0){
                    keep[v] = maxOf(0, c - maxOf(0, (a[v]!! - b.getOrDefault(v, 0))))   // bug: the most changeable is (a-b) instead of (a-keep)
                }
            }
            var same = 0
            for((v, c) in a)
                same += minOf(c, b.getOrDefault(v, 0))
            diff_sum += M - same
        }
        var zeros = keep.values.sum()
        var swaps = diff_sum - (M - zeros)
        println("Case #${it}: $swaps")
    }
}
