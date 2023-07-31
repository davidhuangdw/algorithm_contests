package google_kickstart.`2022`.G
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (rs, rh) = input_ints()
        val ma = (rs+rh) * (rs+rh)
        val (N,) = input_ints()
        val da = (0 until N).map{
            val (x, y) = input_ints()
            x*x + y*y
        }.sorted()
        val (M, ) = input_ints()
        val db = (0 until M).map{
            val (x, y) = input_ints()
            x*x + y*y
        }.sorted()

        fun res(): Pair<Int, Int>{
            val (a, b) = 0 to 0
            if(N == 0 && M == 0)
                return a to b
            if(N > 0 && (M == 0 || da[0] < db[0])){
                var i = 0
                while(i < N){
                    if(da[i] > ma || (M > 0 && da[i] > db[0]))
                        break
                    i++
                }
                return i to 0
            }
            var i = 0
            while(i < M){
                if(db[i] > ma || (N > 0 && db[i] > da[0]))
                    break
                i++
            }
            return 0 to i
        }
        val (a, b) = res()
        println("Case #${it}: $a $b")
    }
}
