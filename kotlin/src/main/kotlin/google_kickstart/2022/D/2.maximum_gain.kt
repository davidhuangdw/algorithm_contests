package google_kickstart.`2022`.D

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    fun pre_sum(arr: List<Long>, k: Int): Pair<List<Long>, List<Long>> {
        val n = arr.size
        val l = mutableListOf(0L)
        val r = mutableListOf(0L)

        for(i in (1..k)){
            l.add(l[i-1] + arr[i-1])
            r.add(r[i-1] + arr[n-i])
        }
        return Pair(l, r)
    }

    fun max_sub(sums: Pair<List<Long>, List<Long>>, k: Int): Long{
        val (lsum, rsum) = sums
        return (0..k).map{ l ->
            lsum[l] + rsum[k-l]
        }.maxOrNull()!!
    }

    val (T, ) = input_ints()
    (1..T).forEach {
        input_ints()
        val A = input_longs()
        input_ints()
        val B = input_longs()
        val (K, ) = input_ints()

        val u = minOf(A.size, K)
        val v = minOf(B.size, K)
        val asums = pre_sum(A, u)
        val bsums = pre_sum(B, v)
        val mx = (K-v..u).map{max_sub(asums,it) + max_sub(bsums, K-it)}.maxOrNull()!!

        println("Case #${it}: $mx")
    }
}
