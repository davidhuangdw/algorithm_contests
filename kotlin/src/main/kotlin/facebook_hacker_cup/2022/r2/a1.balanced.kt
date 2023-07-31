
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val S = split_input()[0]
        val N = S.length
        val Q = input_ints()[0]
        val acc = (1..26).map{ MutableList<Int>(N+1){0} }
        for((i, ch) in S.withIndex()){
            for(j in 0 until 26)
                acc[j][i+1] = acc[j][i] + (if(j == ch -'a') 1 else 0)
        }
        fun counts(i: Int, j: Int) = List(26){k -> acc[k][j+1] - acc[k][i]}
        fun isHalf(half: List<Int>, full: List<Int>): Boolean{
            var diff = 0
            for(k in 0 until 26){
                val a = half[k]
                val b = full[k]
                if(a * 2 == b) continue
                if(a * 2 + 1 == b) {
                    diff++
                    if(diff > 1)
                        return false
                }else
                    return false
            }
            return true
        }
        fun check(l: Int, r: Int): Boolean{
            var k = r - l
            if(k % 2 != 0)
                return false
            k = k/2
            val full = counts(l, r)
            return isHalf(counts(l, l+k-1), full) || isHalf(counts(r-k+1, r), full)
        }
        var c = 0
        for(i in 0 until Q){
            val (l, r) = input_ints()
            val sub = S.substring(l-1 .. r-1)
//            println(sub)
//            println(check(l-1, r-1))
            if(check(l-1, r-1)) c++
        }
        println("Case #${it}: $c")
    }
}
