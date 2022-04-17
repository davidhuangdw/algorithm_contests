package google_code_jam.`2020`.r2

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        var (L, R) = input_longs()
//        var (L, R) = 0L to 0L
        fun set1(ll: Long, rr: Long): String{
            var (L, R) = ll to rr
            var i = 1
            while(true){
                if(i > maxOf(L, R)) break
                if(L >= R) L -= i
                else R-= i
                i += 1
            }
            return "${i-1} $L $R"

        }

        fun set2(ll: Long, rr: Long): String{
            var (L, R) = ll to rr
            fun acc(st: Long, k: Long, step: Int=1) = (st + st+(k-1)*step)*k/2
            fun maxk(st: Long, cap: Long, step: Int=1): Long { // max k that acc(st, k) <= cap
                var (l, r) = 1L to minOf(2000_000_000L, cap)  // bug: r, must >=1 && must not overflow -- minOf 2e9 not 1e9
                while(l <= r){
                    val md = (l+r)/2
                    if(acc(st, md, step) <= cap) l = md+1
                    else r = md -1
                }
                return r
            }

            var cur = 1L
            if(L <= R){
                val k = maxk(cur, R-L)
                R -= acc(cur, k)
                cur += k
                if(L < R && R >= cur){
                    R -= cur
                    cur ++
                }
            }else{
                val k = maxk(cur, L-R)
                L -= acc(cur, k)
                cur += k
            }

            val lk = maxk(cur, L, 2)
            L -= acc(cur, lk, 2)
            val rk = maxk(cur+1, R, 2)
            R -= acc(cur+1, rk, 2)
            val k = cur - 1 + lk + rk
            return "$k $L $R"
        }

//        for(l in 1..5L)
//            for(r in 1..5L){
//                if(set1(l, r) != set2(l, r)){
//                    println("$l $r: ${set1(l, r)} ${set2(l, r)}")
//                }
//            }
        val res = set2(L, R)
        println("Case #${it}: $res")
    }
}
