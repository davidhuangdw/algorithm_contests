
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val S = input_ints()
        fun check(): Boolean {
            if(N > K*2) return false
            val count = HashMap<Int, Int>()
            for(v in S){
                count[v] = count.getOrDefault(v, 0) + 1
            }
            var base = 0
            for((v, c) in count){
                if(c > 2)
                    return false
                if(c == 2) {
                    base++
                    if(base > K)
                        return false
                }
            }
            return true
        }
        val r = if(check()) "YES" else "NO"
        println("Case #${it}: ${r}")
    }
}
