package google_kickstart.`2022`.B
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    fun isPar(x: Long): Boolean{
        val s = x.toString()
        var (l, r)= 0 to s.length-1
        while(l < r){
            if(s[l] != s[r])
                return false
            l ++
            r --
        }
        return true
    }
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (A, ) = input_longs()
        var cnt = 0L
        for(x in 1L..A){
            if(x * x > A)
                break
            if(A % x == 0L){
                if(isPar(x))
                    cnt ++
                if(A/x != x && isPar(A/x))
                    cnt ++
            }
        }
        println("Case #${it}: $cnt")
    }
}
